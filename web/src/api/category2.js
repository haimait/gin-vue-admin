import service from '@/utils/request'

// @Tags Category2
// @Summary 创建分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category2 true "创建分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category2/createCategory2 [post]
export const createCategory2 = (data) => {
  return service({
    url: '/category2/createCategory2',
    method: 'post',
    data
  })
}

// @Tags Category2
// @Summary 删除分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category2 true "删除分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category2/deleteCategory2 [delete]
export const deleteCategory2 = (data) => {
  return service({
    url: '/category2/deleteCategory2',
    method: 'delete',
    data
  })
}

// @Tags Category2
// @Summary 批量删除分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category2/deleteCategory2 [delete]
export const deleteCategory2ByIds = (data) => {
  return service({
    url: '/category2/deleteCategory2ByIds',
    method: 'delete',
    data
  })
}

// @Tags Category2
// @Summary 更新分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category2 true "更新分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category2/updateCategory2 [put]
export const updateCategory2 = (data) => {
  return service({
    url: '/category2/updateCategory2',
    method: 'put',
    data
  })
}

// @Tags Category2
// @Summary 用id查询分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Category2 true "用id查询分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category2/findCategory2 [get]
export const findCategory2 = (params) => {
  return service({
    url: '/category2/findCategory2',
    method: 'get',
    params
  })
}

// @Tags Category2
// @Summary 分页获取分类2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取分类2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category2/getCategory2List [get]
export const getCategory2List = (params) => {
  return service({
    url: '/category2/getCategory2List',
    method: 'get',
    params
  })
}
