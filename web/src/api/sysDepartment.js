import service from '@/utils/request'

// @Tags SysDepartment
// @Summary 创建SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDepartment true "创建SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/createSysDepartment [post]
export const createSysDepartment = (data) => {
  return service({
    url: '/SysDp/createSysDepartment',
    method: 'post',
    data
  })
}

// @Tags SysDepartment
// @Summary 删除SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDepartment true "删除SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SysDp/deleteSysDepartment [delete]
export const deleteSysDepartment = (data) => {
  return service({
    url: '/SysDp/deleteSysDepartment',
    method: 'delete',
    data
  })
}

// @Tags SysDepartment
// @Summary 删除SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SysDp/deleteSysDepartment [delete]
export const deleteSysDepartmentByIds = (data) => {
  return service({
    url: '/SysDp/deleteSysDepartmentByIds',
    method: 'delete',
    data
  })
}

// @Tags SysDepartment
// @Summary 更新SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDepartment true "更新SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SysDp/updateSysDepartment [put]
export const updateSysDepartment = (data) => {
  return service({
    url: '/SysDp/updateSysDepartment',
    method: 'put',
    data
  })
}

// @Tags SysDepartment
// @Summary 用id查询SysDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysDepartment true "用id查询SysDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SysDp/findSysDepartment [get]
export const findSysDepartment = (params) => {
  return service({
    url: '/SysDp/findSysDepartment',
    method: 'get',
    params
  })
}

// @Tags SysDepartment
// @Summary 分页获取SysDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/getSysDepartmentList [get]
export const getSysDepartmentList = (params) => {
  return service({
    url: '/SysDp/getSysDepartmentList',
    method: 'get',
    params
  })
}

// @Tags SysDepartment
// @Summary 分页获取SysDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SysDepartmentSearch true "分页获取SysDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDp/getSysDepartmentTree [get]
export const getSysDepartmentTree = (params) => {
  return service({
    url: '/SysDp/getSysDepartmentTree',
    method: 'get',
    params
  })
}

