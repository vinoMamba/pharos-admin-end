package storage

import (
	"context"

	"github.com/vinoMamba.com/pharos-admin-end/models"
)

// GetRouteList 获取路由列表
func GetRouteList(c context.Context) ([]*models.Menu, error) {
	var menuList []*models.Menu
	err := DB.WithContext(c).Model(models.Menu{}).Where("status = ?", 0).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}

// GetMenuList 获取菜单列表(用于菜单管理)
func GetMenuList(c context.Context) ([]*models.Menu, error) {
	var menuList []*models.Menu
	err := DB.WithContext(c).Model(models.Menu{}).Find(&menuList).Error
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

//	DeleteMenu 删除菜单

func DeleteMenus(c context.Context, idList []int64) error {
	err := DB.WithContext(c).Model(models.Menu{}).Where("id in ?", idList).Update("status", 1).Error
	if err != nil {
		return err
	}
	return nil
}
