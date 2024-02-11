// 自动生成模板Article
package hmArticle

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"gorm.io/datatypes"
)

// article表 结构体  Article
type Article struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"binding:"required"`  //标题 
      Subtitle  string `json:"subtitle" form:"subtitle" gorm:"column:subtitle;comment:副标题;size:255;"`  //副标题 
      Keywords  string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:关键字;size:255;"`  //关键字 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:文章描述;size:1000;"`  //文章描述 
      Author  string `json:"author" form:"author" gorm:"column:author;comment:文章作者;size:60;"`  //文章作者 
      Company  string `json:"company" form:"company" gorm:"column:company;comment:参赛单位;size:60;"`  //参赛单位 
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:11;"`  //手机号 
      Source  string `json:"source" form:"source" gorm:"column:source;comment:文章来源;size:50;"`  //文章来源 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:文章内容;type:text;"`  //文章内容 
      Extend  datatypes.JSON `json:"extend" form:"extend" gorm:"column:extend;comment:附加属性;size:1000;"`  //扩展属性 
      Thumb  string `json:"thumb" form:"thumb" gorm:"column:thumb;comment:封面;size:255;"`  //封面 
      IsLink  *bool `json:"isLink" form:"isLink" gorm:"column:is_link;comment:isLink;size:1;"`  //isLink 
      Link  string `json:"link" form:"link" gorm:"column:link;comment:Link;size:255;"`  //Link 
      Type  string `json:"type" form:"type" gorm:"column:type;type:enum();comment:类型：article/page/video;"`  //类型：article/page/video 
      SortOrder  *int `json:"sortOrder" form:"sortOrder" gorm:"column:sort_order;comment:排序;size:4;"`  //排序 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 1:待审核 2:审核通过;size:1;"`  //状态 1:待审核 2:审核通过 
      Publish  *int `json:"publish" form:"publish" gorm:"column:publish;comment:前台显示 1:是 2:否;size:1;"`  //前台显示 1:是 2:否 
      ViewCount  *int `json:"viewCount" form:"viewCount" gorm:"column:view_count;comment:展示次数;size:11;"`  //展示次数 
      TicketCount  *int `json:"ticketCount" form:"ticketCount" gorm:"column:ticket_count;comment:投票次数;size:11;"`  //投票次数 
      Awards  *int `json:"awards" form:"awards" gorm:"column:awards;comment:获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖;size:1;"`  //获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖 
      Top  *bool `json:"top" form:"top" gorm:"column:top;comment:置顶 1:是 2:否;size:1;"`  //置顶 1:是 2:否 
      AuditorId  *int `json:"auditorId" form:"auditorId" gorm:"column:auditor_id;comment:审核人;size:20;"`  //审核人 
      AuditorFailed  string `json:"auditorFailed" form:"auditorFailed" gorm:"column:auditor_failed;comment:审核结果;size:1000;"`  //审核结果 
      Category_id  *int `json:"category_id" form:"category_id" gorm:"column:category_id;comment:所属分类;size:10;"binding:"required"`  //所属分类 
}


// TableName article表 Article自定义表名 article
func (Article) TableName() string {
  return "article"
}

