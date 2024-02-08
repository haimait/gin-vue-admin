package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Title       string `json:"title" form:"title" `
	Author      string `json:"author" form:"author" `
	Company     string `json:"company" form:"company" `
	Phone       string `json:"phone" form:"phone" `
	IsLink      *bool  `json:"isLink" form:"isLink" `
	Type        string `json:"type" form:"type"`
	Status      *int   `json:"status" form:"status" `
	Publish     *int   `json:"publish" form:"publish" `
	Awards      *int   `json:"awards" form:"awards" `
	Top         *bool  `json:"top" form:"top" `
	AuditorId   *int   `json:"auditorId" form:"auditorId" `
	Category_id *int   `json:"category_id" form:"category_id" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
