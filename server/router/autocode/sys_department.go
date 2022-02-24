package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysDepartmentRouter struct {
}

// InitSysDepartmentRouter 初始化 SysDepartment 路由信息
func (s *SysDepartmentRouter) InitSysDepartmentRouter(Router *gin.RouterGroup) {
	SysDpRouter := Router.Group("SysDp").Use(middleware.OperationRecord())
	SysDpRouterWithoutRecord := Router.Group("SysDp")
	var SysDpApi = v1.ApiGroupApp.AutoCodeApiGroup.SysDepartmentApi
	{
		SysDpRouter.POST("createSysDepartment", SysDpApi.CreateSysDepartment)   // 新建SysDepartment
		SysDpRouter.DELETE("deleteSysDepartment", SysDpApi.DeleteSysDepartment) // 删除SysDepartment
		SysDpRouter.DELETE("deleteSysDepartmentByIds", SysDpApi.DeleteSysDepartmentByIds) // 批量删除SysDepartment
		SysDpRouter.PUT("updateSysDepartment", SysDpApi.UpdateSysDepartment)    // 更新SysDepartment
	}
	{
		SysDpRouterWithoutRecord.GET("findSysDepartment", SysDpApi.FindSysDepartment)        // 根据ID获取SysDepartment
		SysDpRouterWithoutRecord.GET("getSysDepartmentList", SysDpApi.GetSysDepartmentList)  // 获取SysDepartment列表
		SysDpRouterWithoutRecord.GET("getSysDepartmentTree", SysDpApi.GetSysDepartmentTree) // 获取SysDepartmentTree树状列表
	}
}
