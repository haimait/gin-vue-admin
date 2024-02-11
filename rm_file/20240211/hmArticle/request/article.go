package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	"gorm.io/datatypes"
)

type ArticleSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Title  string `json:"title" form:"title" `
                      Author  string `json:"author" form:"author" `
                      Phone  string `json:"phone" form:"phone" `
                      Type  string `json:"type" form:"type"`
                      Status  *int `json:"status" form:"status" `
                      Publish  *int `json:"publish" form:"publish" `
                      Awards  *int `json:"awards" form:"awards" `
                      Top  *bool `json:"top" form:"top" `
                      Category_id  *int `json:"category_id" form:"category_id" `
    request.PageInfo
}
