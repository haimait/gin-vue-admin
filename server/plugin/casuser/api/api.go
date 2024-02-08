package api

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model/casuserRequest"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasUserApi struct{}

// @Tags CasUser
// @Summary 用户注册
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /CasUser/Register [post]
func (p *CasUserApi) Register(c *gin.Context) {
	var cusUser casuserRequest.RegisterCasUser
	err := c.ShouldBindJSON(&cusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//验证入参
	err = utils.Verify(cusUser, utils.RegisterCasUserVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if cusUser.Password != cusUser.RePassword {
		response.FailWithMessage(errors.New("确认密码不一致").Error(), c)
		return
	}
	keyIP := c.ClientIP()

	if err := service.ServiceGroupApp.RegisterService(keyIP, &cusUser); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}

// @Tags CasUser
// @Summary 用户登录
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /CasUser/Login [post]
func (p *CasUserApi) Login(c *gin.Context) {

	if err := service.ServiceGroupApp.LoginService(); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		response.Ok(c)
	}
}
