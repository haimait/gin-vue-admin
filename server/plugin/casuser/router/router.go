package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/api"
	"github.com/gin-gonic/gin"
)

type CasUserRouter struct {
}

func (s *CasUserRouter) InitCasUserRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.CasUserApi
	{
		plugRouter.POST("register", plugApi.Register)
		plugRouter.POST("login", plugApi.Login)
	}
}
