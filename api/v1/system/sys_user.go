package system

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
	"strconv"
)

func (b *BaseApi) Login(c *fiber.Ctx) error {
	var l systemReq.Login
	_ = c.BodyParser(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global.ADTPL_LOG.Error("登录失败！用户名不存在或者密码错误！", zap.Error(err))
			response.FailWithMessage("用户名不存在或密码错误", c)
			return err
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
	return nil
}

func (b *BaseApi) tokenNext(c *fiber.Ctx, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.ADTPL_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.ADTPL_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.ADTPL_CONFIG.System.UseMultipoint {
		response.OkwithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.ADTPL_LOG.Error("设置登录状态失败！", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkwithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.ADTPL_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkwithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func (b *BaseApi) Register(c *fiber.Ctx) error {
	var r systemReq.Register
	_ = c.BodyParser(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{
		Username:    r.Username,
		NickName:    r.NickName,
		Password:    r.Password,
		HeaderImg:   r.HeaderImg,
		AuthorityId: r.AuthorityId,
		Authorities: authorities,
	}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.ADTPL_LOG.Error("注册失败！", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
	return nil
}

func (b *BaseApi) ChangePassword(c *fiber.Ctx) error {
	var user systemReq.ChangePasswordStruct
	_ = c.BodyParser(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	u := &system.SysUser{
		Username: user.Username,
		Password: user.Password,
	}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.ADTPL_LOG.Error("修改失败！", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return err
	} else {
		response.OkWithMessage("修改成功", c)
	}
	return nil
}

func (b *BaseApi) GetUserList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global.ADTPL_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkwithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
	return nil
}

func (b *BaseApi) SetUserAuthority(c *fiber.Ctx) error {
	var sua systemReq.SetUserAuth
	_ = c.BodyParser(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return UserVerifyErr
	}
	userID := utils.GetUserID(c)
	uuid := utils.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		global.ADTPL_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return err
	} else {
		claims := utils.GetUserInfo(c)
		j := &utils.JWT{SigningKey: []byte(global.ADTPL_CONFIG.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			global.ADTPL_LOG.Error("修改失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return err
		} else {
			c.Response().Header.Set("new-token", token)
			c.Response().Header.Set("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response.OkWithMessage("修改成功", c)
		}
	}
	return nil
}

func (b *BaseApi) SetUserAuthorities(c *fiber.Ctx) error {
	var sua systemReq.SetUserAuthorities
	_ = c.BodyParser(&sua)
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		global.ADTPL_LOG.Error("修改失败！", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return err
	} else {
		response.OkWithMessage("修改成功", c)
	}
	return nil
}

func (b *BaseApi) DeleteUser(c *fiber.Ctx) error {
	var reqId request.GetById
	_ = c.BodyParser(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return nil
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.ADTPL_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (b *BaseApi) SetUserInfo(c *fiber.Ctx) error {
	var user system.SysUser
	_ = c.BodyParser(&user)
	user.Username = ""
	user.Password = ""
	user.AuthorityId = ""
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		global.ADTPL_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"userInfo": ReqUser}, "设置成功", c)
	}
	return nil
}

func (b *BaseApi) SetSelfInfo(c *fiber.Ctx) error {
	var user system.SysUser
	_ = c.BodyParser(&user)
	user.Username = ""
	user.Password = ""
	user.AuthorityId = ""
	user.ID = utils.GetUserID(c)
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		global.ADTPL_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"userInfo": ReqUser}, "设置成功", c)
	}
	return nil
}

func (b *BaseApi) GetUserInfo(c *fiber.Ctx) error {
	uuid := utils.GetUserUuid(c)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		global.ADTPL_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"userInfo": ReqUser}, "获取成功", c)
	}
	return nil
}

func (b *BaseApi) ResetPassword(c *fiber.Ctx) error {
	var user system.SysUser
	_ = c.BodyParser(&user)
	if err := userService.ResetPassword(user.ID); err != nil {
		global.ADTPL_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return err
	} else {
		response.OkWithMessage("重置成功", c)
	}
	return nil
}
