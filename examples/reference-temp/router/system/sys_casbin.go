package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router fiber.Router) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithOutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.Post("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithOutRecord.Post("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
