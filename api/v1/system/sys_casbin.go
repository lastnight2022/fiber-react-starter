package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

type CasbinApi struct{}

func (cas *CasbinApi) UpdateCasbin(c *fiber.Ctx) error {
	var cmr request.CasbinInReceive
	_ = c.BodyParser(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.ADTPL_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return err
	} else {
		response.OkWithMessage("更新成功", c)
	}
	return nil
}

func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *fiber.Ctx) error {
	var casbin request.CasbinInReceive
	_ = c.BodyParser(&casbin)
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkwithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
	return nil
}
