package system

import (
	"github.com/gofiber/fiber/v2"
	v1 "server/api/v1"
	"server/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router fiber.Router) (R fiber.Router) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.Post("addBaseMenu", authorityMenuApi.AddBaseMenu)
		menuRouter.Post("addMenuAuthority", authorityMenuApi.AddMenuAuthority)
		menuRouter.Post("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)
		menuRouter.Post("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)
	}
	{
		menuRouterWithoutRecord.Post("getMenu", authorityMenuApi.GetMenu)
		menuRouterWithoutRecord.Post("getMenuList", authorityMenuApi.GetMenuList)
		menuRouterWithoutRecord.Post("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)
		menuRouterWithoutRecord.Post("getMenuAuthority", authorityMenuApi.GetMenuAuthority)
		menuRouterWithoutRecord.Post("getBaseMenuById", authorityMenuApi.GetBaseMenuById)
	}
	return menuRouter
}
