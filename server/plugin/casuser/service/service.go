package service

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model/casuserRequest"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model/casuserResponse"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type CasUserService struct{}

var store = base64Captcha.DefaultMemStore

// RegisterService 用户注册
func (e *CasUserService) RegisterService(keyIP string, in *casuserRequest.RegisterCasUser) (err error) {

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(keyIP)
	if !ok {
		global.BlackCache.Set(keyIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	var captchaBool = store.Verify(in.CaptchaId, in.Captcha, true)

	//开启验证并且没有验证通过 报错误
	if !oc || !captchaBool {
		// 验证码次数+1
		global.BlackCache.Increment(keyIP, 1)
		global.GVA_LOG.Error("验证码错误!", zap.Error(err))
		return errors.New("验证码错误")
	}

	//检查账号是否存在
	var casUser model.Casusers
	ormErr := global.GVA_DB.Where("phone = ?", in.Phone).Or("username = ?", in.Username).First(&casUser).Error
	if ormErr != nil && !errors.Is(ormErr, gorm.ErrRecordNotFound) {
		return err
	}
	if casUser.ID != 0 {
		if casUser.Phone == in.Phone {
			err = errors.New("手机号已存在")
		}
		if casUser.Username == in.Username {
			err = errors.New("用户名已存在")
		}
		return err
	}
	nt := time.Now()
	user := &model.Casusers{
		GVA_MODEL: global.GVA_MODEL{
			CreatedAt: nt,
			UpdatedAt: nt,
		},
		UUID:     uuid.Must(uuid.NewV4()),
		Username: in.Username,
		Nickname: in.Username,
		Phone:    in.Phone,
		Email:    fmt.Sprintf("%s@qq.com", in.Phone),
		Password: utils.BcryptHash(in.Password), //密码加密
	}

	//存储
	err = global.GVA_DB.Create(&user).Error
	return err
}

// LoginService 用户登录
func (e *CasUserService) LoginService(l *casuserRequest.LoginCasUser, c *gin.Context) (err error) {
	keyIP := c.ClientIP()
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(keyIP)
	if !ok {
		global.BlackCache.Set(keyIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		var casUser model.Casusers
		err = global.GVA_DB.Where("phone = ?", l.Phone).Or("username = ?", l.Username).First(&casUser).Error
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(keyIP, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if ok := utils.BcryptCheck(l.Password, casUser.Password); !ok {
			return errors.New("密码错误")
		}
		if casUser.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(keyIP, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		e.TokenNext(c, casUser)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(keyIP, 1)
	response.FailWithMessage("验证码错误", c)

	return nil
}

// TokenNext 登录以后签发jwt
func (e *CasUserService) TokenNext(c *gin.Context, user model.Casusers) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: 0,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(casuserResponse.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(casuserResponse.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(casuserResponse.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
