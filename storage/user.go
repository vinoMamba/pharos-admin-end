package storage

import (
	"context"

	"github.com/vinoMamba.com/pharos-admin-end/models"
)

func GetUserByUsername(c context.Context, username string) (*models.User, error) {
	var u models.User
	if err := DB.WithContext(c).Model(models.User{}).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserById(c context.Context, id int) (*models.User, error) {
	var u models.User
	if err := DB.WithContext(c).Model(models.User{}).Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(c context.Context, u *models.User) error {
	return DB.WithContext(c).Model(models.User{}).Create(u).Error
}

func GetUserListByPage(c context.Context, pageNum, pageSize int) ([]*models.User, int64, error) {
	var userList []*models.User
	var total int64
	err := DB.WithContext(c).Model(models.User{}).Where("is_deleted = ?", 0).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = DB.WithContext(c).Model(models.User{}).Where("is_deleted = ?", 0).Offset((pageNum - 1) * pageSize).Limit(int(pageNum)).Find(&userList).Error
	if err != nil {
		return nil, 0, err
	}
	return userList, total, nil
}

func UpdateUser(c context.Context, u *models.User) error {
	return DB.WithContext(c).Model(models.User{}).Where("id = ?", u.UserId).Updates(u).Error
}

func DeleteUser(c context.Context, id int) error {
	return DB.WithContext(c).Model(models.User{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
