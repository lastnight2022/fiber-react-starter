package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemRes "server/model/system/response"
	"server/utils"
)

type SystemApi struct{}

func (s *SystemApi) GetSystemConfig(c *fiber.Ctx) error {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
	return nil
}

func (s *SystemApi) SetSystemConfig(c *fiber.Ctx) error {
	var sys system.System
	_ = c.BodyParser(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.ADTPL_LOG.Error("设置失败！", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return err
	} else {
		response.OkWithData("设置成功", c)
	}
	return nil
}

func (s *SystemApi) ReloadSystem(c *fiber.Ctx) error {
	err := utils.Reload()
	if err != nil {
		global.ADTPL_LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
		return err
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
	return nil
}

func (s *SystemApi) GetServerInfo(c *fiber.Ctx) error {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkwithDetailed(fiber.Map{"server": server}, "获取成功", c)
	}
	return nil
}
