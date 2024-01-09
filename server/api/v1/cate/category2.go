package cate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cate"
	cateReq "github.com/flipped-aurora/gin-vue-admin/server/model/cate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Category2Api struct {
}

var category2Service = service.ServiceGroupApp.CateServiceGroup.Category2Service

// CreateCategory2 创建分类2
// @Tags Category2
// @Summary 创建分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cate.Category2 true "创建分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category2/createCategory2 [post]
func (category2Api *Category2Api) CreateCategory2(c *gin.Context) {
	var category2 cate.Category2
	err := c.ShouldBindJSON(&category2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category2.CreatedBy = utils.GetUserID(c)
	if err := category2Service.CreateCategory2(&category2); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCategory2 删除分类2
// @Tags Category2
// @Summary 删除分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cate.Category2 true "删除分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category2/deleteCategory2 [delete]
func (category2Api *Category2Api) DeleteCategory2(c *gin.Context) {
	var category2 cate.Category2
	err := c.ShouldBindJSON(&category2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category2.DeletedBy = utils.GetUserID(c)
	if err := category2Service.DeleteCategory2(category2); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCategory2ByIds 批量删除分类2
// @Tags Category2
// @Summary 批量删除分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /category2/deleteCategory2ByIds [delete]
func (category2Api *Category2Api) DeleteCategory2ByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := category2Service.DeleteCategory2ByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCategory2 更新分类2
// @Tags Category2
// @Summary 更新分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cate.Category2 true "更新分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category2/updateCategory2 [put]
func (category2Api *Category2Api) UpdateCategory2(c *gin.Context) {
	var category2 cate.Category2
	err := c.ShouldBindJSON(&category2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category2.UpdatedBy = utils.GetUserID(c)
	if err := category2Service.UpdateCategory2(category2); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCategory2 用id查询分类2
// @Tags Category2
// @Summary 用id查询分类2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cate.Category2 true "用id查询分类2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category2/findCategory2 [get]
func (category2Api *Category2Api) FindCategory2(c *gin.Context) {
	var category2 cate.Category2
	err := c.ShouldBindQuery(&category2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recategory2, err := category2Service.GetCategory2(category2.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory2": recategory2}, c)
	}
}

// GetCategory2List 分页获取分类2列表
// @Tags Category2
// @Summary 分页获取分类2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cateReq.Category2Search true "分页获取分类2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category2/getCategory2List [get]
func (category2Api *Category2Api) GetCategory2List(c *gin.Context) {
	var pageInfo cateReq.Category2Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := category2Service.GetCategory2InfoList(pageInfo); err != nil {
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
