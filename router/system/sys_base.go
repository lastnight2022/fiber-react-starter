package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router fiber.Router) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.Post("login", baseApi.Login)
		baseRouter.Post("captcha", baseApi.Captcha)
	}
}
