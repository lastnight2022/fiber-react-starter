package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router fiber.Router) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.Post("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)
		operationRecordRouter.Delete("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecordByIds)
		operationRecordRouter.Delete("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds)
		operationRecordRouter.Get("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)
		operationRecordRouter.Get("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)
	}
}
