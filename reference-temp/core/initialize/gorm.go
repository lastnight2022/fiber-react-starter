package initialize

import (
	"gorm.io/gorm"
	"server/global"
)

func Gorm() *gorm.DB {
	switch global.ADTPL_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
