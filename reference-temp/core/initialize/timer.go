package initialize

import (
	"fmt"
	"server/config"
	"server/global"
	"server/utils"
)

func Timer() {
	if global.ADTPL_CONFIG.Timer.Start {
		for i := range global.ADTPL_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.ADTPL_Timer.AddTaskByFunc("ClearDB", global.ADTPL_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.ADTPL_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.ADTPL_CONFIG.Timer.Detail[i])
		}
	}
}
