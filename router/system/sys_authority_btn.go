package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
)

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router fiber.Router) {
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouterWithoutRecord.Post("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.Post("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouterWithoutRecord.Post("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn)
	}
}
