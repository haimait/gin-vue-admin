package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CategorySearch struct {
	cate.Category
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
