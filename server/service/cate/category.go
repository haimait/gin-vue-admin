package cate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cate"
	cateReq "github.com/flipped-aurora/gin-vue-admin/server/model/cate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type CategoryService struct {
}

// CreateCategory 创建分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) CreateCategory(category *cate.Category) (err error) {
	err = global.GVA_DB.Create(category).Error
	return err
}

// DeleteCategory 删除分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategory(category cate.Category) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cate.Category{}).Where("id = ?", category.ID).Update("deleted_by", category.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&category).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCategoryByIds 批量删除分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategoryByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cate.Category{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cate.Category{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCategory 更新分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) UpdateCategory(category cate.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategory(id uint) (category cate.Category, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategoryInfoList(info cateReq.CategorySearch) (list []cate.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cate.Category{})
	var categorys []cate.Category
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Enable != nil {
		db = db.Where("enable = ?", info.Enable)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["sort"] = true
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

	err = db.Find(&categorys).Error
	return categorys, total, err
}
