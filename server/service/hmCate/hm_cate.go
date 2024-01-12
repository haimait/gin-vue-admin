package hmCate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hmCate"
    hmCateReq "github.com/flipped-aurora/gin-vue-admin/server/model/hmCate/request"
    "gorm.io/gorm"
)

type HmCategoryService struct {
}

// CreateHmCategory 创建分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService) CreateHmCategory(Category *hmCate.HmCategory) (err error) {
	err = global.GVA_DB.Create(Category).Error
	return err
}

// DeleteHmCategory 删除分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService)DeleteHmCategory(id string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hmCate.HmCategory{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hmCate.HmCategory{},"id = ?",id).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteHmCategoryByIds 批量删除分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService)DeleteHmCategoryByIds(ids []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hmCate.HmCategory{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids).Delete(&hmCate.HmCategory{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateHmCategory 更新分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService)UpdateHmCategory(Category hmCate.HmCategory) (err error) {
	err = global.GVA_DB.Save(&Category).Error
	return err
}

// GetHmCategory 根据id获取分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService)GetHmCategory(id string) (Category hmCate.HmCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Category).Error
	return
}

// GetHmCategoryInfoList 分页获取分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (CategoryService *HmCategoryService)GetHmCategoryInfoList(info hmCateReq.HmCategorySearch) (list []hmCate.HmCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hmCate.HmCategory{})
    var Categorys []hmCate.HmCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Level != nil {
        db = db.Where("level = ?",info.Level)
    }
    if info.ParentId != nil {
        db = db.Where("parent_id = ?",info.ParentId)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.IsShow != nil {
        db = db.Where("is_show = ?",info.IsShow)
    }
    if info.Type != "" {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
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
	
	err = db.Find(&Categorys).Error
	return  Categorys, total, err
}
