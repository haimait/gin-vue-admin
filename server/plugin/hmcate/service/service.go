package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	hmCate "github.com/flipped-aurora/gin-vue-admin/server/plugin/hmCate/model"
	hmCateReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/hmCate/model/request"
	"gorm.io/gorm"
)

type HmcateService struct{}

func (hmCateService *HmcateService) PlugService() (err error) {
	// 写你的业务逻辑
	return nil
}

type HmCategoryService struct {
}

// CreateHmCategory 创建分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (hmCateService *HmcateService) CreateHmCategory(Category *hmCate.HmCategory) (err error) {
	err = global.GVA_DB.Create(Category).Error
	return err
}

// DeleteHmCategory 删除分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (hmCateService *HmcateService) DeleteHmCategory(id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&hmCate.HmCategory{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&hmCate.HmCategory{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteHmCategoryByIds 批量删除分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (hmCateService *HmcateService) DeleteHmCategoryByIds(ids []string, deleted_by uint) (err error) {
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
func (hmCateService *HmcateService) UpdateHmCategory(Category hmCate.HmCategory) (err error) {
	err = global.GVA_DB.Save(&Category).Error
	return err
}

// GetHmCategory 根据id获取分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (hmCateService *HmcateService) GetHmCategory(id string) (Category hmCate.HmCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Category).Error
	return
}

// GetHmCategoryInfoList 分页获取分类管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (hmCateService *HmcateService) GetHmCategoryInfoList(info hmCateReq.HmCategorySearch) (list []hmCate.HmCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hmCate.HmCategory{})
	var Categorys []hmCate.HmCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Level != nil {
		db = db.Where("level = ?", info.Level)
	}
	if info.ParentId != nil {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.IsShow != nil {
		db = db.Where("is_show = ?", info.IsShow)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
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

	err = db.Find(&Categorys).Error
	return Categorys, total, err
}

//@author: haima
//@function: GetHmCategoryTree
//@description: 获取路由分页
//@return: list interface{}, total int64,err error

func (hmCateService *HmcateService) GetHmCategoryTree(pageInfo *hmCateReq.HmCategorySearch) (list interface{}, total int64, err error) {
	var menuList []hmCate.HmCategory
	treeMap, err := hmCateService.getBaseHmCateMenuTreeMap(pageInfo)
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = hmCateService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

//@author: haima
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func (hmCateService *HmcateService) getBaseChildrenList(menu *hmCate.HmCategory, treeMap map[int][]hmCate.HmCategory) (err error) {
	menu.Children = treeMap[int(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = hmCateService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: haima
//@function: getBaseHmCateMenuTreeMap
//@description: 获取路由总树map
//@return: treeMap map[string][]system.SysBaseMenu, err error

func (hmCateService *HmcateService) getBaseHmCateMenuTreeMap(pageInfo *hmCateReq.HmCategorySearch) (treeMap map[int][]hmCate.HmCategory, err error) {
	var allMenus []hmCate.HmCategory
	treeMap = make(map[int][]hmCate.HmCategory)
	db := global.GVA_DB.Order("sort")
	if pageInfo.Type != "" {
		db = db.Where("type = ?", pageInfo.Type)
	}
	err = db.Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[*v.ParentId] = append(treeMap[*v.ParentId], v)
	}
	return treeMap, err
}
