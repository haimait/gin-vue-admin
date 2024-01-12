package hmCate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hmCate"
    hmCateReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmCate/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type HmCategoryApi struct {
}

var CategoryService = service.ServiceGroupApp.HmCateServiceGroup.HmCategoryService


// CreateHmCategory 创建分类管理
// @Tags HmCategory
// @Summary 创建分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmCate.HmCategory true "创建分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Category/createHmCategory [post]
func (CategoryApi *HmCategoryApi) CreateHmCategory(c *gin.Context) {
	var Category hmCate.HmCategory
	err := c.ShouldBindJSON(&Category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    Category.CreatedBy = utils.GetUserID(c)

	if err := CategoryService.CreateHmCategory(&Category); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHmCategory 删除分类管理
// @Tags HmCategory
// @Summary 删除分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmCate.HmCategory true "删除分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Category/deleteHmCategory [delete]
func (CategoryApi *HmCategoryApi) DeleteHmCategory(c *gin.Context) {
	id := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := CategoryService.DeleteHmCategory(id,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHmCategoryByIds 批量删除分类管理
// @Tags HmCategory
// @Summary 批量删除分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Category/deleteHmCategoryByIds [delete]
func (CategoryApi *HmCategoryApi) DeleteHmCategoryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
    userID := utils.GetUserID(c)
	if err := CategoryService.DeleteHmCategoryByIds(ids,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHmCategory 更新分类管理
// @Tags HmCategory
// @Summary 更新分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hmCate.HmCategory true "更新分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Category/updateHmCategory [put]
func (CategoryApi *HmCategoryApi) UpdateHmCategory(c *gin.Context) {
	var Category hmCate.HmCategory
	err := c.ShouldBindJSON(&Category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    Category.UpdatedBy = utils.GetUserID(c)

	if err := CategoryService.UpdateHmCategory(Category); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHmCategory 用id查询分类管理
// @Tags HmCategory
// @Summary 用id查询分类管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmCate.HmCategory true "用id查询分类管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Category/findHmCategory [get]
func (CategoryApi *HmCategoryApi) FindHmCategory(c *gin.Context) {
	id := c.Query("ID")
	if reCategory, err := CategoryService.GetHmCategory(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCategory": reCategory}, c)
	}
}

// GetHmCategoryList 分页获取分类管理列表
// @Tags HmCategory
// @Summary 分页获取分类管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hmCateReq.HmCategorySearch true "分页获取分类管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Category/getHmCategoryList [get]
func (CategoryApi *HmCategoryApi) GetHmCategoryList(c *gin.Context) {
	var pageInfo hmCateReq.HmCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := CategoryService.GetHmCategoryInfoList(pageInfo); err != nil {
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
