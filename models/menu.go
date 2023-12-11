package models

import "time"

type Menu struct {
	MenuId           int64     `json:"menuId" gorm:"column:id;primaryKey;not null"`
	ParentId         string    `json:"parentId" gorm:"column:parent_id;default:null"`
	ParentIds        string    `json:"parentIds" gorm:"column:parent_ids;default:null"`
	MenuName         string    `json:"menuName" gorm:"column:menu_name;not null"`
	RoutePath        string    `json:"routePath" gorm:"column:route_path;not null"`
	RouteName        string    `json:"routeName" gorm:"column:route_name;not null"`
	Redirect         string    `json:"redirect" gorm:"column:redirect;default:null"`
	Component        string    `json:"component" gorm:"column:component;default:null"`
	Permission       string    `json:"permission" gorm:"column:permission;default:null"`
	Type             int       `json:"type" gorm:"column:type;not null"` //  0：系统 1：目录 2：菜单 3：按钮
	Affix            int       `json:"affix" gorm:"column:affix;default:null"`
	Icon             string    `json:"icon" gorm:"column:icon;default:null"`
	Sort             int       `json:"sort" gorm:"sort;default:null"`
	Endpoint         int       `json:"endpoint" gorm:"end_point;default:0"` //  0：管理后台 1：移动端
	Status           int       `json:"status" gorm:"status;default:0"`      // 状态 0：正常 1：冻结
	HideMenu         int       `json:"hideMenu" gorm:"hide_menu;default:0"`
	HideChildrenMenu int       `json:"hideChildrenMenu" gorm:"hide_children_menu;default:0"`
	IsDeleted        int       `json:"isDeleted" gorm:"is_deleted;default:0"` // 0：正常 1：删除
	CreateBy         string    `json:"createBy" gorm:"create_by;default:null"`
	CreateTime       time.Time `json:"createTime" gorm:"create_time;default:null"`
	UpdateBy         string    `json:"updateBy" gorm:"update_by;default:null"`
	UpdateTime       time.Time `json:"updateTime" gorm:"update_time;default:null"`
}

func (a Menu) TableName() string {
	return "sys_menu"
}
