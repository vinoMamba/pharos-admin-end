package storage

import (
	"context"

	"github.com/spf13/cast"
	"github.com/vinoMamba.com/pharos-admin-end/models"
	"github.com/vinoMamba.com/pharos-admin-end/utils"
)

// 根据角色Id和 menuId 的数组，批量创建数据
func CreateRoleMenus(c context.Context, roleId int64, menuIds []string) error {
	var list []models.RoleMenu
	for _, menuId := range menuIds {
		list = append(list, models.RoleMenu{
			Id:     utils.GetSnowflakeIdInt64(),
			RoleId: roleId,
			MenuId: cast.ToInt64(menuId),
		})
	}
	return DB.WithContext(c).Create(&list).Error
}

// 根据角色Id数组，删除角色菜单关联表中的数据
func DeleteRoleMenuByRoleId(c context.Context, roleIds []int64) error {
	return DB.WithContext(c).Where("role_id IN ?", roleIds).Delete(&models.RoleMenu{}).Error
}

// 根据角色Id，查询角色菜单Id数组
func GetMenuIdsByRoleId(c context.Context, roleId int64) ([]int64, error) {
	var list []models.RoleMenu
	err := DB.WithContext(c).Where("role_id = ?", roleId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var menuIds []int64
	for _, item := range list {
		menuIds = append(menuIds, item.MenuId)
	}
	return menuIds, nil
}
