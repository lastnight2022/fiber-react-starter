package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

type SystemApiApi struct{}

func (s *SystemApiApi) CreateApi(c *fiber.Ctx) error {
	var api system.SysApi
	_ = c.BodyParser(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := apiService.CreateApi(api); err != nil {
		global.ADTPL_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return err
	} else {
		response.OkWithMessage("创建成功", c)
	}
	return nil
}

func (s *SystemApiApi) DeleteApi(c *fiber.Ctx) error {
	var api system.SysApi
	_ = c.BodyParser(&api)
	if err := utils.Verify(api.ADTPL_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := apiService.DeleteApi(api); err != nil {
		global.ADTPL_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (s *SystemApiApi) GetApiList(c *fiber.Ctx) error {
	var pageInfo systemReq.SearchApiParams
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, list, total := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.OkWithMessage("获取失败！", c)
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

func (s *SystemApiApi) GetApiById(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	err, api := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkWithData(systemRes.SysAPIResponse{
			Api: api,
		}, c)
	}
	return nil
}

func (s *SystemApiApi) UpdateApi(c *fiber.Ctx) error {
	var api system.SysApi
	_ = c.BodyParser(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err := apiService.UpdateApi(api); err != nil {
		global.ADTPL_LOG.Error("修改失败！", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return err
	} else {
		response.OkWithMessage("修改成功", c)
	}
	return nil
}

func (s *SystemApiApi) GetAllApis(c *fiber.Ctx) error {
	if err, apis := apiService.GetAllApis(); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "获取成功", c)
	}
	return nil
}

func (s *SystemApiApi) DeleteApisByIds(c *fiber.Ctx) error {
	var ids request.IdsReq
	_ = c.BodyParser(&ids)
	if err := apiService.DeleteApisByIds(ids); err != nil {
		global.ADTPL_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}
