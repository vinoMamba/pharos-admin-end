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
