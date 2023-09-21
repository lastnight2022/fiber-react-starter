package system

import "server/global"

type JwtBlacklist struct {
	global.ADTPL_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
