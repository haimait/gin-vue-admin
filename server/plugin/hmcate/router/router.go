package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/hmcate/api"
	"github.com/gin-gonic/gin"
)

type HmCateRouter struct {
}

func (s *HmCateRouter) InitHmCateRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	CategoryRouterWithoutRecord := Router
	plugApi := api.ApiGroupApp.HmcateApi
	{
		//CategoryRouter.POST("routerName", plugApi.ApiName)
		CategoryRouter.POST("createHmCategory", plugApi.CreateHmCategory)             // 新建分类管理
		CategoryRouter.DELETE("deleteHmCategory", plugApi.DeleteHmCategory)           // 删除分类管理
		CategoryRouter.DELETE("deleteHmCategoryByIds", plugApi.DeleteHmCategoryByIds) // 批量删除分类管理
		CategoryRouter.PUT("updateHmCategory", plugApi.UpdateHmCategory)
	}
	{
		CategoryRouterWithoutRecord.GET("findHmCategory", plugApi.FindHmCategory)       // 根据ID获取分类管理
		CategoryRouterWithoutRecord.GET("getHmCategoryList", plugApi.GetHmCategoryList) // 获取分类管理列表
		CategoryRouterWithoutRecord.GET("getHmCategoryTree", plugApi.GetHmCategoryTree) // 获取分类管理列表树

	}
}
