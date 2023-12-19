package models

import "time"

type RoleMenu struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;not null"`
	RoleId     int64     `json:"roleId" gorm:"column:role_id;default:null"`
	MenuId     int64     `json:"menuId" gorm:"column:menu_id;default:null"`
	CreateBy   string    `json:"createBy" gorm:"create_by;default:null"`
	CreateTime time.Time `json:"createTime" gorm:"create_time;default:null"`
}

func (a RoleMenu) TableName() string {
	return "sys_role_menu"
}
