package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDepartmentApi struct {
}

var SysDpService = service.ServiceGroupApp.AutoCodeServiceGroup.SysDepartmentService

// CreateSysDepartment 创建SysDepartment
// @Tags SysDepartment
// @Summary 创建SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysDepartment true "创建SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/createSysDepartment [post]
func (SysDpApi *SysDepartmentApi) CreateSysDepartment(c *gin.Context) {
	var SysDp autocode.SysDepartment
	_ = c.ShouldBindJSON(&SysDp)
	if err := SysDpService.CreateSysDepartment(SysDp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysDepartment 删除SysDepartment
// @Tags SysDepartment
// @Summary 删除SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysDepartment true "删除SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SysDp/deleteSysDepartment [delete]
func (SysDpApi *SysDepartmentApi) DeleteSysDepartment(c *gin.Context) {
	var SysDp autocode.SysDepartment
	_ = c.ShouldBindJSON(&SysDp)
	if err := SysDpService.DeleteSysDepartment(SysDp); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("删除失败:%s", err.Error()), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysDepartmentByIds 批量删除SysDepartment
// @Tags SysDepartment
// @Summary 批量删除SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /SysDp/deleteSysDepartmentByIds [delete]
func (SysDpApi *SysDepartmentApi) DeleteSysDepartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := SysDpService.DeleteSysDepartmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysDepartment 更新SysDepartment
// @Tags SysDepartment
// @Summary 更新SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SysDepartment true "更新SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SysDp/updateSysDepartment [put]
func (SysDpApi *SysDepartmentApi) UpdateSysDepartment(c *gin.Context) {
	var SysDp autocode.SysDepartment
	_ = c.ShouldBindJSON(&SysDp)
	if err := SysDpService.UpdateSysDepartment(SysDp); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysDepartment 用id查询SysDepartment
// @Tags SysDepartment
// @Summary 用id查询SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SysDepartment true "用id查询SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SysDp/findSysDepartment [get]
func (SysDpApi *SysDepartmentApi) FindSysDepartment(c *gin.Context) {
	var SysDp autocode.SysDepartment
	_ = c.ShouldBindQuery(&SysDp)
	if err, reSysDp := SysDpService.GetSysDepartment(SysDp.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSysDp": reSysDp}, c)
	}
}

// GetSysDepartmentList 分页获取SysDepartment列表
// @Tags SysDepartment
// @Summary 分页获取SysDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysDepartmentSearch true "分页获取SysDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/getSysDepartmentList [get]
func (SysDpApi *SysDepartmentApi) GetSysDepartmentList(c *gin.Context) {
	var pageInfo autocodeReq.SysDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := SysDpService.GetSysDepartmentInfoList(pageInfo); err != nil {
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

// GetSysDepartmentTree 获取SysDepartment所有数据按树状排序
// @Tags SysDepartment
// @Summary 分页获取SysDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysDepartmentSearch true "分页获取SysDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/getSysDepartmentTree [get]
func (SysDpApi *SysDepartmentApi) GetSysDepartmentTree(c *gin.Context) {
	var pageInfo autocodeReq.SysDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, err := SysDpService.GetDepartmentTree(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(len(list)),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
