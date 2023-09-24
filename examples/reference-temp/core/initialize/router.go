package initialize

import (
	"github.com/gofiber/fiber/v2"
	"server/global"
	"server/middleware"
	"server/router"
)

func Routers(app *fiber.App) {

	app.Static(global.ADTPL_CONFIG.Local.Prefix, global.ADTPL_CONFIG.Local.RootPath)

	global.ADTPL_LOG.Info("use middleware logger")

	global.ADTPL_LOG.Info("use middleware cors")

	systemRouter := router.RouterGroupApp.System

	PublicGroup := app.Group("")
	{
		PublicGroup.Get("/health", func(c *fiber.Ctx) error {
			c.Status(200).JSON("ok")
			return nil
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := app.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitBaseRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
	}

	global.ADTPL_LOG.Info("router register success")
}
