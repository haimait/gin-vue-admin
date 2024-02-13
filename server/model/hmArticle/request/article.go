package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Title      string `json:"title" form:"title" `
	Author     string `json:"author" form:"author" `
	Company    string `json:"company" form:"company" `
	Phone      string `json:"phone" form:"phone" `
	IsLink     int    `json:"isLink" form:"isLink" `
	Type       string `json:"type" form:"type"`
	Status     int    `json:"status" form:"status" `
	Publish    int    `json:"publish" form:"publish" `
	IsAwards   int    `json:"isAwards" form:"isAwards" ` //是否获奖 1.是 2.否
	Awards     int    `json:"awards" form:"awards" `
	Top        int    `json:"top" form:"top" `
	AuditorId  int    `json:"auditorId" form:"auditorId" `
	CategoryId int    `json:"categoryId" form:"categoryId" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
