package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	password = password + "hjxz"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func CheckHashPassword(hasPwd, pwd string) bool {
	pwd = pwd + "hjxz"
	bytes, err := base64.StdEncoding.DecodeString(hasPwd)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bytes, []byte(pwd))
	return err == nil
}
