package main

import (
	"go.uber.org/zap"
	"server/core"
	"server/core/initialize"
	"server/global"
)

func main() {
	global.ADTPL_VP = core.Viper()
	global.ADTPL_LOG = core.Zap()
	zap.ReplaceGlobals(global.ADTPL_LOG)
	global.ADTPL_DB = initialize.Gorm()
	global.ADTPL_Pool = initialize.NewPool()
	global.ADTPL_Mongo = initialize.NewMongo()
	initialize.Timer()
	initialize.DBList()
	core.RunServer()
}
