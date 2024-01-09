package cate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Category2Router struct {
}

// InitCategory2Router 初始化 分类2 路由信息
func (s *Category2Router) InitCategory2Router(Router *gin.RouterGroup) {
	category2Router := Router.Group("category2").Use(middleware.OperationRecord())
	category2RouterWithoutRecord := Router.Group("category2")
	var category2Api = v1.ApiGroupApp.CateApiGroup.Category2Api
	{
		category2Router.POST("createCategory2", category2Api.CreateCategory2)             // 新建分类2
		category2Router.DELETE("deleteCategory2", category2Api.DeleteCategory2)           // 删除分类2
		category2Router.DELETE("deleteCategory2ByIds", category2Api.DeleteCategory2ByIds) // 批量删除分类2
		category2Router.PUT("updateCategory2", category2Api.UpdateCategory2)              // 更新分类2
	}
	{
		category2RouterWithoutRecord.GET("findCategory2", category2Api.FindCategory2)       // 根据ID获取分类2
		category2RouterWithoutRecord.GET("getCategory2List", category2Api.GetCategory2List) // 获取分类2列表
	}
}
