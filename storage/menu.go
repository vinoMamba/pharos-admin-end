package storage

import (
	"context"

	"github.com/vinoMamba.com/pharos-admin-end/models"
)

// GetRouteList 获取路由列表
func GetRouteList(c context.Context) ([]*models.Menu, error) {
	var menuList []*models.Menu
	err := DB.WithContext(c).Model(models.Menu{}).Where("status = ?", 0).Where("type != ?", 3).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}

// GetMenuList 获取菜单列表(用于菜单管理)
func GetMenuList(c context.Context) ([]*models.Menu, error) {
	var menuList []*models.Menu
	err := DB.WithContext(c).Model(models.Menu{}).Where("is_deleted", 0).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}

// SaveMenu 保存菜单
func SaveMenu(c context.Context, menu *models.Menu) error {
	err := DB.WithContext(c).Model(models.Menu{}).Create(menu).Error
	if err != nil {
		return err
	}
	return nil
}

// updateMenu 更新菜单
func UpdateMenu(c context.Context, menu *models.Menu) error {
	return DB.WithContext(c).Where("id = ?", menu.MenuId).Updates(menu).Error
}

//	DeleteMenu 删除菜单

func DeleteMenus(c context.Context, idList []int64) error {
	err := DB.WithContext(c).Model(models.Menu{}).Where("id IN ?", idList).Update("is_deleted", 1).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMenuById 根据ID获取菜单
func GetMenuById(c context.Context, id int64) (*models.Menu, error) {
	var menu models.Menu
	err := DB.WithContext(c).Model(models.Menu{}).Where("id = ?", id).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// 获取permission code
func GetPermCode(c context.Context) ([]string, error) {
	perCodeList := make([]string, 0)
	err := DB.WithContext(c).Model(models.Menu{}).Select("permission").Where("permission != ''").Find(&perCodeList).Error
	if err != nil {
		return nil, err
	}
	return perCodeList, nil
}
