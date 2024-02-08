package service

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/casuser/model/casuserRequest"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"github.com/mojocn/base64Captcha"
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
func (e *CasUserService) LoginService() (err error) {
	// 写你的业务逻辑
	return nil
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
