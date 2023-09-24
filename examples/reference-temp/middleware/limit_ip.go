package middleware

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"server/global"
	"server/model/common/response"
	"time"
)

type LimitConfig struct {
	GenerationKey func(c *fiber.Ctx) string
	CheckOrMark   func(key string, expire int, limit int) error
	Expire        int
	Limit         int
}

func (l LimitConfig) LimitWithTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.Status(http.StatusOK).JSON(fiber.Map{"code": response.ERROR, "msg": err})
			return err
		} else {
			c.Next()
		}
		return nil
	}
}

func DefaultLimit() fiber.Handler {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.ADTPL_CONFIG.System.LimitTimeIP,
		Limit:         global.ADTPL_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}

func DefaultGenerationKey(c *fiber.Ctx) string {
	return "ADTPL_Limit" + c.IP()
}

func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.ADTPL_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.ADTPL_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// 次数
		if times, err := global.ADTPL_REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.ADTPL_REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.ADTPL_REDIS.Incr(context.Background(), key).Err()
			}
		}
	}
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	if global.ADTPL_REDIS == nil {
		return err
	}
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.ADTPL_LOG.Error("limit", zap.Error(err))
	}
	return err
}
