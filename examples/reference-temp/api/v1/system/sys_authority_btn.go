package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system/request"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

func (a *AuthorityBtnApi) GetAuthorityBtn(c *fiber.Ctx) error {
	var req request.SysAuthorityBtnReq
	_ = c.BodyParser(&req)
	if err, res := authorityBtnService.GetAuthorityBtn(req); err != nil {
		global.ADTPL_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return err
	} else {
		response.OkwithDetailed(res, "查询成功", c)
	}
	return nil
}

func (a *AuthorityBtnApi) SetAuthorityBtn(c *fiber.Ctx) error {
	var req request.SysAuthorityBtnReq
	_ = c.BodyParser(&req)
	if err := authorityBtnService.SetAuthorityBtn(req); err != nil {
		global.ADTPL_LOG.Error("分配失败", zap.Error(err))
		response.FailWithMessage("分配失败", c)
		return err
	} else {
		response.OkWithMessage("分配成功", c)
	}
	return nil
}

func (a *AuthorityBtnApi) CanRemoveAuthorityBtn(c *fiber.Ctx) error {
	id := c.Query("id")
	if err := authorityBtnService.CanRemoveAuthorityBtn(id); err != nil {
		global.ADTPL_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}
