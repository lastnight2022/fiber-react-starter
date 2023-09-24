package middleware

import (
	"github.com/gofiber/fiber/v2"
	"server/config"
	"server/global"
)

func Cors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := c.Method()
		origin := c.Get("Origin")
		c.Response().Header.Set("Access-Control-Allow-Origin", origin)
		c.Response().Header.Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Response().Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Response().Header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			return nil
		}

		c.Next()
		return nil
	}
}

func CorsByRules() fiber.Handler {
	if global.ADTPL_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *fiber.Ctx) error {
		whitelist := checkCors(c.Get("origin"))
		if whitelist != nil {
			c.Response().Header.Set("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Response().Header.Set("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Response().Header.Set("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Response().Header.Set("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
			}
		}

		if whitelist == nil && global.ADTPL_CONFIG.Cors.Mode == "strict-whitelist" && !(c.Method() == "GET" && c.Path() == "/health") {
			return nil
		} else {
			if c.Method() == "OPTIONS" {
				return nil
			}
		}
		c.Next()
		return nil
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.ADTPL_CONFIG.Cors.Whitelist {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
