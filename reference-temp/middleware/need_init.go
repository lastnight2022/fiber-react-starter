package middleware

import (
	"github.com/gofiber/fiber/v2"
	"server/global"
	"server/model/common/response"
)

func NeedInit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if global.ADTPL_DB == nil {
			response.OkwithDetailed(fiber.Map{
				"needInit": true,
			}, "前往初始化数据库", c)
		} else {
			c.Next()
		}
		return nil
	}
}
