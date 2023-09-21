package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"server/global"
)

func Redis() {
	redisCfg := global.ADTPL_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ADTPL_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.ADTPL_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.ADTPL_REDIS = client
	}
}
