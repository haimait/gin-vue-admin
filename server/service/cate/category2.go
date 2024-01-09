package cate

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cate"
	cateReq "github.com/flipped-aurora/gin-vue-admin/server/model/cate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type Category2Service struct {
}

// CreateCategory2 创建分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) CreateCategory2(category2 *cate.Category2) (err error) {
	err = global.GVA_DB.Create(category2).Error
	return err
}

// DeleteCategory2 删除分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) DeleteCategory2(category2 cate.Category2) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cate.Category2{}).Where("id = ?", category2.ID).Update("deleted_by", category2.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&category2).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCategory2ByIds 批量删除分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) DeleteCategory2ByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cate.Category2{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cate.Category2{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCategory2 更新分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) UpdateCategory2(category2 cate.Category2) (err error) {
	err = global.GVA_DB.Save(&category2).Error
	return err
}

// GetCategory2 根据id获取分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) GetCategory2(id uint) (category2 cate.Category2, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category2).Error
	return
}

// GetCategory2InfoList 分页获取分类2记录
// Author [piexlmax](https://github.com/piexlmax)
func (category2Service *Category2Service) GetCategory2InfoList(info cateReq.Category2Search) (list []cate.Category2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cate.Category2{})
	var category2s []cate.Category2
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
	orderMap["order"] = true
	if orderMap[info.Sort] {
		OrderStr = fmt.Sprintf("`%s`", info.Sort)
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&category2s).Error
	return category2s, total, err
}
