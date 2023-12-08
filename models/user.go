package models

type User struct {
	UserId   int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	RealName string `json:"realName" gorm:"column:real_name"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avatar"`
}

func (a User) TableName() string {
	return "sys_user"
}
