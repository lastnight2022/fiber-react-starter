package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"server/global"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.ADTPL_CONFIG.System.DbType {
	case "mysql":
		logZap = global.ADTPL_CONFIG.Mysql.LogZap
		break
	default:
		logZap = global.ADTPL_CONFIG.Mysql.LogZap
	}
	if logZap {
		global.ADTPL_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
