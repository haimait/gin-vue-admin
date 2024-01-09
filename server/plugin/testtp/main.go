package testtp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/testtp/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/testtp/router"
	"github.com/gin-gonic/gin"
)

type TestTPPlugin struct {
}

func CreateTestTPPlug(maxNum int) *TestTPPlugin {
	global.GlobalConfig.MaxNum = maxNum

	return &TestTPPlugin{}
}

func (*TestTPPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitTestTPRouter(group)
}

func (*TestTPPlugin) RouterPath() string {
	return "testTP"
}
