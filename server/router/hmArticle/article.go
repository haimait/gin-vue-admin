package hmArticle

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

// InitArticleRouter 初始化 article表 路由信息
func (s *ArticleRouter) InitArticleRouter(PrivateGroup *gin.RouterGroup, PublicGroup *gin.RouterGroup) {
	articleRouter := PrivateGroup.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := PublicGroup.Group("article")
	articleRouterWithoutView := PublicGroup.Group("viewArticle")
	var articleApi = v1.ApiGroupApp.ArticleApiGroup.ArticleApi
	{
		articleRouter.POST("createArticle", articleApi.CreateArticle)             // 新建article表
		articleRouter.DELETE("deleteArticle", articleApi.DeleteArticle)           // 删除article表
		articleRouter.DELETE("deleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除article表
		articleRouter.PUT("updateArticle", articleApi.UpdateArticle)              // 更新article表
	}
	{
		articleRouterWithoutView.POST("createArticle", articleApi.CreateArticle)    // 新建article表
		articleRouterWithoutView.PUT("updateArticle", articleApi.UpdateArticle)     // 更新article表
		articleRouterWithoutRecord.GET("findArticle", articleApi.FindArticle)       // 根据ID获取article表
		articleRouterWithoutRecord.GET("getArticleList", articleApi.GetArticleList) // 获取article表列表
	}
}