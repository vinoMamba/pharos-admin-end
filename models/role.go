package models

type Role struct {
	RoleId    int64  `json:"roleId" gorm:"column:id;primaryKey;not null"`
	RoleName  string `json:"roleName" gorm:"column:role_name"`
	RoleCode  string `json:"roleCode" gorm:"column:role_code"`
	Remark    string `json:"remark" gorm:"column:remark"`
	Status    int    `json:"status" gorm:"status"`        // 状态 0：正常 1：冻结
	IsDeleted int    `json:"isDeleted" gorm:"is_deleted"` // 0：正常 1：删除
}

func (a Role) TableName() string {
	return "sys_role"
}
