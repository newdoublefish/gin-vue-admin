import service from '@/utils/request'

// @Tags AutoPosition
// @Summary 创建AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoPosition true "创建AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Position/createAutoPosition [post]
export const createAutoPosition = (data) => {
  return service({
    url: '/Position/createAutoPosition',
    method: 'post',
    data
  })
}

// @Tags AutoPosition
// @Summary 删除AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoPosition true "删除AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Position/deleteAutoPosition [delete]
export const deleteAutoPosition = (data) => {
  return service({
    url: '/Position/deleteAutoPosition',
    method: 'delete',
    data
  })
}

// @Tags AutoPosition
// @Summary 删除AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Position/deleteAutoPosition [delete]
export const deleteAutoPositionByIds = (data) => {
  return service({
    url: '/Position/deleteAutoPositionByIds',
    method: 'delete',
    data
  })
}

// @Tags AutoPosition
// @Summary 更新AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoPosition true "更新AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Position/updateAutoPosition [put]
export const updateAutoPosition = (data) => {
  return service({
    url: '/Position/updateAutoPosition',
    method: 'put',
    data
  })
}

// @Tags AutoPosition
// @Summary 用id查询AutoPosition
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AutoPosition true "用id查询AutoPosition"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Position/findAutoPosition [get]
export const findAutoPosition = (params) => {
  return service({
    url: '/Position/findAutoPosition',
    method: 'get',
    params
  })
}

// @Tags AutoPosition
// @Summary 分页获取AutoPosition列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AutoPosition列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Position/getAutoPositionList [get]
export const getAutoPositionList = (params) => {
  return service({
    url: '/Position/getAutoPositionList',
    method: 'get',
    params
  })
}
