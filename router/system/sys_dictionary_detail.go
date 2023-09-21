package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router fiber.Router) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.Post("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)
		dictionaryDetailRouter.Delete("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail)
		dictionaryDetailRouter.Put("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)
	}
	{
		dictionaryDetailRouterWithoutRecord.Get("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)
		dictionaryDetailRouterWithoutRecord.Get("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList)
	}
}
