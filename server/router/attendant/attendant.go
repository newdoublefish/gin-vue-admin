package attendant

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AttendantRouter struct {

}

func (s *AttendantRouter) InitAutoPositionRouter(Router *gin.RouterGroup) {
	PositionRouterWithoutRecord := Router.Group("attendant")
	var api = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		PositionRouterWithoutRecord.POST("getUserAttendant", api.GetUserAttendant)   // 获取上下班时间
	}
}
