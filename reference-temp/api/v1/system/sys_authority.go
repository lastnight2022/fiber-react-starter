package system

import (
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

func (a *AuthorityApi) CreateAuthority(c *fiber.Ctx) error {
	var authority system.SysAuthority
	_ = c.BodyParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.ADTPL_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
		return err
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkwithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
	return nil
}

func (a *AuthorityApi) CopyAuthority(c *fiber.Ctx) error {
	var copyInfo systemRes.SysAuthorityCopyResponse
	_ = c.BodyParser(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
		global.ADTPL_LOG.Error("拷贝失败", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
	return nil
}

func (a *AuthorityApi) DeleteAuthority(c *fiber.Ctx) error {
	var authority system.SysAuthority
	_ = c.BodyParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil {
		global.ADTPL_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除成功"+err.Error(), c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (a *AuthorityApi) UpdateAuthority(c *fiber.Ctx) error {
	var auth system.SysAuthority
	_ = c.BodyParser(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		global.ADTPL_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
	return nil
}

func (a *AuthorityApi) GetAuthorityList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.ADTPL_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return err
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

func (a *AuthorityApi) SetDataAuthority(c *fiber.Ctx) error {
	var auth system.SysAuthority
	_ = c.BodyParser(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := authorityService.SetDataAuthority(auth); err != nil {
		global.ADTPL_LOG.Error("设置失败！", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return err
	} else {
		response.OkWithMessage("设置成功", c)
	}
	return nil
}
