// 自动生成模板Category2
package cate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 分类2 结构体  Category2
type Category2 struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:分类名称;size:100;"`             //分类名称
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:191;"`         //备注
	Order     *int   `json:"order"  gorm:"column:order;comment:排序;size:10;"`                         //排序
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:是否启用 1正常 2冻结;size:1;"` //是否启用
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 分类2 Category2自定义表名 category12
func (Category2) TableName() string {
	return "category12"
}
