package storage

import (
	"context"

	"github.com/vinoMamba.com/pharos-admin-end/models"
)

func GetMenuList(c context.Context) ([]*models.Menu, error) {
	var menuList []*models.Menu
	// 从数据库获取menu 列表
	err := DB.WithContext(c).Model(models.Menu{}).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}
