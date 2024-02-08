package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
	"time"
)

type Casusers struct {
	global.GVA_MODEL
	UUID         uuid.UUID `json:"uuid" gorm:"uniqueIndex;comment:用户UUID"`                    // 用户UUID
	Username     string    `json:"username" gorm:"type:varchar(255);uniqueIndex;comment:用户名"` //
	Nickname     string    `json:"nickname" gorm:"type:varchar(255);comment:昵称"`              //
	Email        string    `json:"email" gorm:"type:varchar(255);uniqueIndex;comment:Email"`
	Phone        string    `json:"phone" gorm:"type:varchar(11);uniqueIndex;comment:手机号"`         //
	Enable       int       `json:"enable" gorm:"type:int(1);default:1;comment:用户是否被冻结 1正常 2冻结"`   //用户是否被冻结 1正常 2冻结
	Password     string    `json:"-"  gorm:"column:password;type:varchar(255);comment:用户登录密码"`    // 用户登录密码
	SignInCount  int64     `json:"signInCount" gorm:"type:int(11);comment:登录次数"`                  //
	LastSignInAt time.Time `json:"lastSignInAt" gorm:"type:datetime;default:null;comment:最近登录时间"` //
	LastSignInIp string    `json:"lastSignInIp" gorm:"type:varchar(255);comment:最近登录IP"`          //
	Description  string    `json:"description" gorm:"type:varchar(11);comment:用户说明"`              //
}

//ConfirmPassword string         `json:"confirmPassword" gorm:"-"` //确认登陆密码                   //

func (Casusers) TableName() string {
	return "casusers"
}
