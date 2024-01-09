package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/testtp/api"
	"github.com/gin-gonic/gin"
)

type TestTPRouter struct {
}

func (s *TestTPRouter) InitTestTPRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.TestTPApi
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
