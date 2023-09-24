package global

import (
	"github.com/alitto/pond"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"server/config"
	"server/utils/timer"
	"sync"
)

var (
	ADTPL_DB                  *gorm.DB
	ADTPL_DBList              map[string]*gorm.DB
	ADTPL_Mongo               *MongoDB
	ADTPL_REDIS               *redis.Client
	ADTPL_LOG                 *zap.Logger
	ADTPL_CONFIG              config.Server
	ADTPL_VP                  *viper.Viper
	ADTPL_Pool                *pond.WorkerPool
	ADTPL_Timer               timer.Timer = timer.NewTimerTask()
	ADTPL_Concurrency_Control             = &singleflight.Group{}
	BlackCache                local_cache.Cache
	lock                      sync.RWMutex
)

func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return ADTPL_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := ADTPL_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
