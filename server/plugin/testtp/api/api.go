package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/testtp/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestTPApi struct{}

// @Tags TestTP
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /testTP/routerName [post]
func (p *TestTPApi) ApiName(c *gin.Context) {

	if err := service.ServiceGroupApp.PlugService(); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		response.OkWithData("成功", c)
	}
}
