package hmcate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/hmcate/router"
	"github.com/gin-gonic/gin"
)

type HmcatePlugin struct {
}

func CreateHmcatePlug() *HmcatePlugin {
	return &HmcatePlugin{}
}

func (*HmcatePlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitHmCateRouter(group)
}

func (*HmcatePlugin) RouterPath() string {
	return "category"
}
