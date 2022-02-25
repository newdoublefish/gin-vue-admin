package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AutoPositionRouter struct {
}

// InitAutoPositionRouter 初始化 AutoPosition 路由信息
func (s *AutoPositionRouter) InitAutoPositionRouter(Router *gin.RouterGroup) {
	PositionRouter := Router.Group("Position").Use(middleware.OperationRecord())
	PositionRouterWithoutRecord := Router.Group("Position")
	var PositionApi = v1.ApiGroupApp.AutoCodeApiGroup.AutoPositionApi
	{
		PositionRouter.POST("createAutoPosition", PositionApi.CreateAutoPosition)   // 新建AutoPosition
		PositionRouter.DELETE("deleteAutoPosition", PositionApi.DeleteAutoPosition) // 删除AutoPosition
		PositionRouter.DELETE("deleteAutoPositionByIds", PositionApi.DeleteAutoPositionByIds) // 批量删除AutoPosition
		PositionRouter.PUT("updateAutoPosition", PositionApi.UpdateAutoPosition)    // 更新AutoPosition
	}
	{
		PositionRouterWithoutRecord.GET("findAutoPosition", PositionApi.FindAutoPosition)        // 根据ID获取AutoPosition
		PositionRouterWithoutRecord.GET("getAutoPositionList", PositionApi.GetAutoPositionList)  // 获取AutoPosition列表
	}
}
