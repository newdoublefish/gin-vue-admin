package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AutoPositionApi struct {
}

var PositionService = service.ServiceGroupApp.AutoCodeServiceGroup.AutoPositionService


// CreateAutoPosition 创建AutoPosition
// @Tags AutoPosition
// @Summary 创建AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoPosition true "创建AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Position/createAutoPosition [post]
func (PositionApi *AutoPositionApi) CreateAutoPosition(c *gin.Context) {
	var Position autocode.AutoPosition
	_ = c.ShouldBindJSON(&Position)
	if err := PositionService.CreateAutoPosition(Position); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAutoPosition 删除AutoPosition
// @Tags AutoPosition
// @Summary 删除AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoPosition true "删除AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Position/deleteAutoPosition [delete]
func (PositionApi *AutoPositionApi) DeleteAutoPosition(c *gin.Context) {
	var Position autocode.AutoPosition
	_ = c.ShouldBindJSON(&Position)
	if err := PositionService.DeleteAutoPosition(Position); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAutoPositionByIds 批量删除AutoPosition
// @Tags AutoPosition
// @Summary 批量删除AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Position/deleteAutoPositionByIds [delete]
func (PositionApi *AutoPositionApi) DeleteAutoPositionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := PositionService.DeleteAutoPositionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAutoPosition 更新AutoPosition
// @Tags AutoPosition
// @Summary 更新AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.AutoPosition true "更新AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Position/updateAutoPosition [put]
func (PositionApi *AutoPositionApi) UpdateAutoPosition(c *gin.Context) {
	var Position autocode.AutoPosition
	_ = c.ShouldBindJSON(&Position)
	if err := PositionService.UpdateAutoPosition(Position); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAutoPosition 用id查询AutoPosition
// @Tags AutoPosition
// @Summary 用id查询AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.AutoPosition true "用id查询AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Position/findAutoPosition [get]
func (PositionApi *AutoPositionApi) FindAutoPosition(c *gin.Context) {
	var Position autocode.AutoPosition
	_ = c.ShouldBindQuery(&Position)
	if err, rePosition := PositionService.GetAutoPosition(Position.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePosition": rePosition}, c)
	}
}

// GetAutoPositionList 分页获取AutoPosition列表
// @Tags AutoPosition
// @Summary 分页获取AutoPosition列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AutoPositionSearch true "分页获取AutoPosition列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Position/getAutoPositionList [get]
func (PositionApi *AutoPositionApi) GetAutoPositionList(c *gin.Context) {
	var pageInfo autocodeReq.AutoPositionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := PositionService.GetAutoPositionInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
