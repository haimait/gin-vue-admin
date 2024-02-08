package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	HmCategoryTypeArticle = "article"
	HmCategoryTypeGood    = "goods"
)

// 分类管理 结构体  HmCategory
type HmCategory struct {
	global.GVA_MODEL
	Level     *int         `json:"level" form:"level" gorm:"column:level;comment:分类层级;size:1;"`                                                              //分类层级
	ParentId  *int         `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父级id;size:20;"`                                                   //父级id
	Name      string       `json:"name" form:"name" gorm:"column:name;comment:分类名称;size:191;"`                                                               //分类名称
	IsShow    *bool        `json:"isShow" form:"isShow" gorm:"column:is_show;comment:是否显示 0:隐藏 1:显示;size:1;"`                                                //是否显示
	ImgPath   string       `json:"imgPath" form:"imgPath" gorm:"column:img_path;default:https://qmplusimg.henrongyi.top/gvalogo.png;comment:图片地址;size:191;"` //图片地址
	Type      string       `json:"type" form:"type" gorm:"column:type;type:enum('article','goods');comment:类型：article/goods;"`                               //类型
	Sort      *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:1;"`                                                                   //排序
	Extend    string       `json:"extend" form:"extend" gorm:"column:extend;comment:扩展字段;size:1000;"`                                                        //扩展字段
	CreatedBy uint         `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint         `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint         `gorm:"column:deleted_by;comment:删除者"`
	Children  []HmCategory `json:"children" gorm:"-"`
}

// TableName 分类管理 HmCategory自定义表名 hm_cate
func (HmCategory) TableName() string {
	return "hm_cate"
}
