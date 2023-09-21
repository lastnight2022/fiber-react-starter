package core

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"server/core/initialize"
	"server/global"
	"server/service/system"
	"time"
)

func RunServer() {
	if global.ADTPL_CONFIG.System.UseMultipoint || global.ADTPL_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.ADTPL_DB != nil {
		system.LoadAll()
	}

	app := fiber.New()

	initialize.Routers(app)

	address := fmt.Sprintf(":%d", global.ADTPL_CONFIG.System.Addr)
	time.Sleep(10 * time.Microsecond)
	global.ADTPL_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`"Welcome to admin template server,address: %s"`, address)
	global.ADTPL_LOG.Error(app.Listen(address).Error())
}
