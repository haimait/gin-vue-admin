package hmCate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HmCategoryRouter struct {
}

// InitHmCategoryRouter 初始化 分类管理 路由信息
func (s *HmCategoryRouter) InitHmCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("Category").Use(middleware.OperationRecord())
	CategoryRouterWithoutRecord := Router.Group("Category")
	var CategoryApi = v1.ApiGroupApp.HmCateApiGroup.HmCategoryApi
	{
		CategoryRouter.POST("createHmCategory", CategoryApi.CreateHmCategory)   // 新建分类管理
		CategoryRouter.DELETE("deleteHmCategory", CategoryApi.DeleteHmCategory) // 删除分类管理
		CategoryRouter.DELETE("deleteHmCategoryByIds", CategoryApi.DeleteHmCategoryByIds) // 批量删除分类管理
		CategoryRouter.PUT("updateHmCategory", CategoryApi.UpdateHmCategory)    // 更新分类管理
	}
	{
		CategoryRouterWithoutRecord.GET("findHmCategory", CategoryApi.FindHmCategory)        // 根据ID获取分类管理
		CategoryRouterWithoutRecord.GET("getHmCategoryList", CategoryApi.GetHmCategoryList)  // 获取分类管理列表
	}
}
