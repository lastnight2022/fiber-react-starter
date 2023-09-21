package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router fiber.Router) {
	apiRouter := Router.Group("api", middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi

	{
		apiRouter.Post("createApi", apiRouterApi.CreateApi)
		apiRouter.Post("deleteApi", apiRouterApi.DeleteApi)
		apiRouter.Post("getApiById", apiRouterApi.GetApiById)
		apiRouter.Post("updateApi", apiRouterApi.UpdateApi)
		apiRouter.Delete("deleteApisByIds", apiRouterApi.DeleteApisByIds)
	}

	{
		apiRouterWithoutRecord.Post("getAllApis", apiRouterApi.GetAllApis)
		apiRouterWithoutRecord.Post("getApiList", apiRouterApi.GetApiList)
	}
}
