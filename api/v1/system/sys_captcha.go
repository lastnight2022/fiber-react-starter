package system

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	systemRes "server/model/system/response"
	"server/utils/captcha"
)

var store = captcha.NewDefaultRedisStore()

type BaseApi struct{}

func (b *BaseApi) Captcha(c *fiber.Ctx) error {
	driver := base64Captcha.NewDriverDigit(
		global.ADTPL_CONFIG.Captcha.ImgHeight,
		global.ADTPL_CONFIG.Captcha.ImgWidth,
		global.ADTPL_CONFIG.Captcha.KeyLong,
		0.7,
		80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c.Context()))
	if id, b64s, err := cp.Generate(); err != nil {
		global.ADTPL_LOG.Error("验证码获取失败！", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return err
	} else {
		response.OkwithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.ADTPL_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
	}
	return nil
}
