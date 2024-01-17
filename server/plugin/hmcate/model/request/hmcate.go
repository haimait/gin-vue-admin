package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HmCategorySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Level          *int       `json:"level" form:"level" `
	ParentId       *int       `json:"parentId" form:"parentId" `
	Name           string     `json:"name" form:"name" `
	IsShow         *bool      `json:"isShow" form:"isShow" `
	Type           string     `json:"type" form:"type"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
