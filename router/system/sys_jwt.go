package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router fiber.Router) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.Post("jsonInBlacklist", jwtApi.JsonInBlacklist)
	}
}
