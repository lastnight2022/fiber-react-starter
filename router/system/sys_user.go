package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router fiber.Router) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.Post("register", baseApi.Register)
		userRouter.Post("changePassword", baseApi.ChangePassword)
		userRouter.Post("setUserAuthority", baseApi.SetUserAuthority)
		userRouter.Delete("deleteUser", baseApi.DeleteUser)
		userRouter.Put("setUserInfo", baseApi.SetUserInfo)
		userRouter.Put("setSelfInfo", baseApi.SetSelfInfo)
		userRouter.Post("setUserAuthorities", baseApi.SetUserAuthorities)
		userRouter.Post("resetPassword", baseApi.ResetPassword)
	}
	{
		userRouterWithoutRecord.Post("getUserList", baseApi.GetUserList)
		userRouterWithoutRecord.Get("getUserInfo", baseApi.GetUserInfo)
	}
}
