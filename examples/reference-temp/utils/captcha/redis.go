package captcha

import (
	"context"
	"go.uber.org/zap"
	"server/global"
	"time"

	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

func (rs *RedisStore) Set(id string, value string) {
	err := global.ADTPL_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.ADTPL_LOG.Error("RedisStoreGetError!", zap.Error(err))
	}
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.ADTPL_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.ADTPL_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.ADTPL_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.ADTPL_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}
