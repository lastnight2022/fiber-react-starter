package initialize

import (
	"gorm.io/gorm"
	"server/global"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.ADTPL_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Dbname] = GormMysqlByConfig(info)
		default:
			continue
		}
	}
	if sysDB, ok := dbMap[sys]; ok {
		global.ADTPL_DB = sysDB
	}
	global.ADTPL_DBList = dbMap
}
