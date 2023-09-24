package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
	"server/service"
	"server/utils"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("x-token")
		if token == "" {
			response.FailWithDetailed(fiber.Map{"reload": true}, "未登录或非法登录", c)
			return errors.New("未登录或非法登录")
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(fiber.Map{"reload": true}, "您的账户异地登录或令牌失效", c)
			return errors.New("您的账户异地登录或令牌失效")
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(fiber.Map{"reload": true}, "授权已过期", c)
				return errors.New("授权已过期")
			}
			response.FailWithDetailed(fiber.Map{"reload": true}, err.Error(), c)
			return err
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.ADTPL_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Response().Header.Set("new-token", newToken)
			c.Response().Header.Set("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.ADTPL_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.ADTPL_LOG.Error("get redis jwt failed", zap.Error(err))
				} else {
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Locals("claims", claims)
		return nil
	}
}
