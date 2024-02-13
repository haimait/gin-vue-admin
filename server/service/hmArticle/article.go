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
func (articleService *ArticleService) DeleteArticle(ID string) (err error) {
	err = global.GVA_DB.Delete(&hmArticle.Article{}, "id = ?", ID).Error
	return err
}

// DeleteArticleByIds 批量删除article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hmArticle.Article{}, "id in ?", IDs).Error
	return err
}

// UpdateArticle 更新article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) UpdateArticle(article hmArticle.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据ID获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticle(ID string) (article hmArticle.Article, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&article).Error
	return
}

// GetArticleInfoList 分页获取article表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticleInfoList(info hmArticleReq.ArticleSearch) (list []hmArticle.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hmArticle.Article{}).Preload("Category")
	var articles []hmArticle.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Author != "" {
		db = db.Where("author = ?", info.Author)
	}
	if info.Company != "" {
		db = db.Where("company = ?", info.Company)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.IsLink != 0 {
		db = db.Where("is_link = ?", info.IsLink)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.Publish != 0 {
		db = db.Where("publish = ?", info.Publish)
	}
	if info.IsAwards != 0 {
		if info.IsAwards == 1 {
			db = db.Where("awards != ?", 5) //中奖
		} else {
			db = db.Where("awards = ?", 5) //未中奖
		}
	}
	if info.Awards != 0 {
		db = db.Where("awards = ?", info.Awards)
	}
	if info.Top != 0 {
		db = db.Where("top = ?", info.Top)
	}
	if info.AuditorId != 0 {
		db = db.Where("auditor_id = ?", info.AuditorId)
	}
	if info.CategoryId != 0 {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["sort_order"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&articles).Error
	return articles, total, err
}
