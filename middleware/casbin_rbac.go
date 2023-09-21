package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/global"
	"server/model/common/response"
	"server/service"
	"server/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器

func CasbinHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		waitUse, _ := utils.GetClaims(c)
		obj := c.Path()
		act := c.Method()
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		success, _ := e.Enforce(sub, obj, act)
		if global.ADTPL_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(fiber.Map{}, "权限不足", c)
			return errors.New("权限不足")
		}
		return nil
	}
}
