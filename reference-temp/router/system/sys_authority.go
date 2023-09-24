package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router fiber.Router) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.Post("createAuthority", authorityApi.CreateAuthority)
		authorityRouter.Post("deleteAuthority", authorityApi.DeleteAuthority)
		authorityRouter.Put("updateAuthority", authorityApi.UpdateAuthority)
		authorityRouter.Post("copyAuthority", authorityApi.CopyAuthority)
		authorityRouter.Post("setDataAuthority", authorityApi.SetDataAuthority)
	}
	{
		authorityRouterWithoutRecord.Post("getAuthorityList", authorityApi.GetAuthorityList)
	}
}
