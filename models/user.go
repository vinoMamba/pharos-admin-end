package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password"`
}

func (a User) TableName() string {
	return "sys_user"
}
