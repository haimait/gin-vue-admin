package hmArticle

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle"
    hmArticleReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmArticle/request"
)

type ArticleService struct {
}

// CreateArticle 创建article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article *hmArticle.Article) (err error) {
	err = global.GVA_DB.Create(article).Error
	return err
}

// DeleteArticle 删除article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticle(ID string) (err error) {
	err = global.GVA_DB.Delete(&hmArticle.Article{},"id = ?",ID).Error
	return err
}

// DeleteArticleByIds 批量删除article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hmArticle.Article{},"id in ?",IDs).Error
	return err
}

// UpdateArticle 更新article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)UpdateArticle(article hmArticle.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据ID获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticle(ID string) (article hmArticle.Article, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&article).Error
	return
}

// GetArticleInfoList 分页获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticleInfoList(info hmArticleReq.ArticleSearch) (list []hmArticle.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hmArticle.Article{})
    var articles []hmArticle.Article
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Author != "" {
        db = db.Where("author = ?",info.Author)
    }
    if info.Phone != "" {
        db = db.Where("phone = ?",info.Phone)
    }
    if info.Type != "" {
        db = db.Where("type = ?",info.Type)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
    if info.Publish != nil {
        db = db.Where("publish = ?",info.Publish)
    }
    if info.Awards != nil {
        db = db.Where("awards = ?",info.Awards)
    }
    if info.Top != nil {
        db = db.Where("top = ?",info.Top)
    }
    if info.Category_id != nil {
        db = db.Where("category_id = ?",info.Category_id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&articles).Error
	return  articles, total, err
}
