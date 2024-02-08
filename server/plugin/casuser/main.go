package casuser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/router"
	"github.com/gin-gonic/gin"
)

type CasUserPlugin struct {
}

func CreateCasUserPlug() *CasUserPlugin {
	global.GVA_DB.AutoMigrate(model.Casusers{})
	return &CasUserPlugin{}
}

func (*CasUserPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitCasUserRouter(group)
}

func (*CasUserPlugin) RouterPath() string {
	return "casUser"
}
