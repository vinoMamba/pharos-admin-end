package storage

import (
	"context"

	"github.com/vinoMamba.com/pharos-admin-end/models"
)

// 角色分页表
func GetRoleListByPage(c context.Context, pageNum, pageSize int) ([]*models.Role, int64, error) {
	var roleList []*models.Role
	var total int64
	err := DB.WithContext(c).Model(models.Role{}).Where("is_deleted = ?", 0).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = DB.WithContext(c).Model(models.Role{}).Where("is_deleted = ?", 0).Offset((pageNum - 1) * pageSize).Limit(int(pageNum)).Find(&roleList).Error
	if err != nil {
		return nil, 0, err
	}
	return roleList, total, nil
}

// 全部角色
func GetAllRoleList(c context.Context) ([]*models.Role, error) {
	var roleList []*models.Role
	err := DB.WithContext(c).Model(models.Role{}).Where("is_deleted = ?", 0).Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	return roleList, nil
}

// 添加角色
func AddRole(c context.Context, role *models.Role) error {
	err := DB.WithContext(c).Model(models.Role{}).Create(role).Error
	if err != nil {
		return err
	}
	return nil
}

// 查询角色
func GetRoleById(c context.Context, roleId int64) (*models.Role, error) {
	var role models.Role
	err := DB.WithContext(c).Model(models.Role{}).Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// 更新角色
func UpdateRole(c context.Context, role *models.Role) error {
	return DB.WithContext(c).Save(role).Error
}

// 删除角色(更新is_delete字段)
func DeleteRole(c context.Context, roleIds []int64) error {
	err := DB.WithContext(c).Model(models.Role{}).Where("id IN ?", roleIds).Update("is_deleted", 1).Error
	if err != nil {
		return err
	}
	return nil
}

// 角色详情
func GetRoleDetail(c context.Context, roleId int64) (*models.Role, error) {
	var role models.Role
	err := DB.WithContext(c).Model(models.Role{}).Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
