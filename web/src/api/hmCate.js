import service from '@/utils/request'

// @Tags HmCategory
// @Summary 创建分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HmCategory true "创建分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Category/createHmCategory [post]
export const createHmCategory = (data) => {
  return service({
    url: '/Category/createHmCategory',
    method: 'post',
    data
  })
}

// @Tags HmCategory
// @Summary 删除分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HmCategory true "删除分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Category/deleteHmCategory [delete]
export const deleteHmCategory = (params) => {
  return service({
    url: '/Category/deleteHmCategory',
    method: 'delete',
    params
  })
}

// @Tags HmCategory
// @Summary 批量删除分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Category/deleteHmCategory [delete]
export const deleteHmCategoryByIds = (params) => {
  return service({
    url: '/Category/deleteHmCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags HmCategory
// @Summary 更新分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HmCategory true "更新分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Category/updateHmCategory [put]
export const updateHmCategory = (data) => {
  return service({
    url: '/Category/updateHmCategory',
    method: 'put',
    data
  })
}

// @Tags HmCategory
// @Summary 用id查询分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HmCategory true "用id查询分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Category/findHmCategory [get]
export const findHmCategory = (params) => {
  return service({
    url: '/Category/findHmCategory',
    method: 'get',
    params
  })
}

// @Tags HmCategory
// @Summary 分页获取分类管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取分类管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Category/getHmCategoryList [get]
export const getHmCategoryList = (params) => {
  return service({
    url: '/Category/getHmCategoryList',
    method: 'get',
    params
  })
}
