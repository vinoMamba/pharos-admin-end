package models

import "time"

type User struct {
	UserId    int64     `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	TenantId  string    `json:"tenantId" gorm:"column:tenant_id"`
	RealName  string    `json:"realName" gorm:"column:real_name"`
	JobNumber string    `json:"jobNumber" gorm:"column:job_number"`
	Mobile    string    `json:"mobile" gorm:"column:mobile"`
	Email     string    `json:"email" gorm:"column:email"`
	Avatar    string    `json:"avatar"`
	HireDate  string    `json:"hireDate" gorm:"column:hire_date"`
	Status    int       `json:"status" gorm:"column:status"`
	LeaveTime time.Time `json:"leaveTime" gorm:"column:leave_time"`
	IsAdmin   int       `json:"isAdmin" gorm:"column:is_admin"`
}

func (a User) TableName() string {
	return "sys_user"
}
