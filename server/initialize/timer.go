package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}

	// 半夜统计一次, 时间 0 0 10,14,16 * * ?
	//global.GVA_Timer.AddTaskByFunc("statistic", "0 39 13 2 4 ? 2022", func() {
	//	// global.GVA_Timer.AddTaskByFunc("statistic", "0 12 * * ?", func() {
	//	//service.Statistic(5)
	//	fmt.Println("execute ------------------")
	//})
	us := userService
	// 每日同步时间
	global.GVA_Timer.AddTaskByFunc("sync", "0 0 12,18,22 * * ?", func() {
		us.SyncUsersFromAttendantSystem()
	})

}
