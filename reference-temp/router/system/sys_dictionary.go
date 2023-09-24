package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router fiber.Router) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.Post("createSysDictionary", sysDictionaryApi.CreateSysDictionary)
		sysDictionaryRouter.Delete("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary)
		sysDictionaryRouter.Put("updateSysDictionary", sysDictionaryApi.UpdatesysDictionary)
	}
	{
		sysDictionaryRouterWithoutRecord.Get("findSysDictionary", sysDictionaryApi.FindSysDictionary)
		sysDictionaryRouterWithoutRecord.Get("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList)
	}
}
