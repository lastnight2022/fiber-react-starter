package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	"server/utils"
)

type OperationRecordApi struct{}

func (s *OperationRecordApi) CreateSysOperationRecord(c *fiber.Ctx) error {
	var sysOperationRecord system.SysOperationRecord
	_ = c.BodyParser(&sysOperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.ADTPL_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return err
	} else {
		response.FailWithMessage("创建成功", c)
	}
	return nil
}

func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.BodyParser(&IDS)
	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.ADTPL_LOG.Error("批量删除失败！", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return err
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
	return nil
}

func (s *OperationRecordApi) FindSysOperationRecord(c *fiber.Ctx) error {
	var sysOperationRecord system.SysOperationRecord
	_ = c.BodyParser(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return err
	}
	if err, resysOperationRecord := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.ADTPL_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkwithDetailed(fiber.Map{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
	return nil
}

func (s *OperationRecordApi) GetSysOperationRecordList(c *fiber.Ctx) error {
	var pageInfo systemReq.SysOperationRecordSearch
	_ = c.BodyParser(&pageInfo)
	if err, list, total := operationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
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
