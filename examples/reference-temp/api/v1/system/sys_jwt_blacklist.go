package system

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
)

type JwtApi struct{}

func (j *JwtApi) JsonInBlacklist(c *fiber.Ctx) error {
	token := c.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.ADTPL_LOG.Error("jwt作废失败！", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return err
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
	return nil
}
