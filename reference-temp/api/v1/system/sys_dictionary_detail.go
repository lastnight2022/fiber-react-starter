package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
	"server/model/system/request"
	"server/utils"
)

type DictionaryDetailApi struct{}

func (s *DictionaryDetailApi) CreateSysDictionaryDetail(c *fiber.Ctx) error {
	var detail system.SysDictionaryDetail
	_ = c.BodyParser(&detail)
	if err := dictionaryDetailService.CreateSysDictionaryDetail(detail); err != nil {
		global.ADTPL_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return err
	} else {
		response.OkWithMessage("创建成功", c)
	}
	return nil
}

func (s *DictionaryDetailApi) DeleteSysDictionaryDetail(c *fiber.Ctx) error {
	var detail system.SysDictionaryDetail
	_ = c.BodyParser(&detail)
	if err := dictionaryDetailService.DeleteSysDictionaryDetail(detail); err != nil {
		global.ADTPL_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return err
	} else {
		response.OkWithMessage("删除成功", c)
	}
	return nil
}

func (s *DictionaryDetailApi) UpdateSysDictionaryDetail(c *fiber.Ctx) error {
	var detail system.SysDictionaryDetail
	_ = c.BodyParser(&detail)
	if err := dictionaryDetailService.UpdateSysDictionaryDetail(&detail); err != nil {
		global.ADTPL_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return err
	} else {
		response.OkWithMessage("更新成功", c)
	}
	return nil
}

func (s *DictionaryDetailApi) FindSysDictionaryDetail(c *fiber.Ctx) error {
	var detail system.SysDictionaryDetail
	_ = c.BodyParser(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, resysDictionaryDetail := dictionaryDetailService.GetSysDictionaryDetail(detail.ID); err != nil {
		global.ADTPL_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return err
	} else {
		response.OkwithDetailed(fiber.Map{"resysDictionaryDetail": resysDictionaryDetail}, "查询成功", c)
	}
	return nil
}

func (s *DictionaryDetailApi) GetSysDictionaryDetailList(c *fiber.Ctx) error {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.BodyParser(&pageInfo)
	if err, list, total := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
		global.ADTPL_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
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
