package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router fiber.Router) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.Post("getSystemConfig", systemApi.GetSystemConfig)
		sysRouter.Post("setSystemConfig", systemApi.SetSystemConfig)
		sysRouter.Post("getServerInfo", systemApi.GetServerInfo)
		sysRouter.Post("reloadSystem", systemApi.ReloadSystem)
	}
}
