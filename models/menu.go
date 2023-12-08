package models

type Menu struct {
	MenuId           string `json:"menuId" gorm:"column:id"`
	MenuName         string `json:"menuName" gorm:"column:menu_name"`
	ParentId         string `json:"parentId" gorm:"column:parent_id"`
	RoutePath        string `json:"routePath" gorm:"column:route_path"`
	RouteName        string `json:"routeName" gorm:"column:route_name"`
	Redirect         string `json:"redirect" gorm:"column:redirect"`
	Component        string `json:"component" gorm:"column:component"`
	Type             int    `json:"type" gorm:"column:type"`
	Affix            int    `json:"affix" gorm:"column:affix"`
	Icon             string `json:"icon" gorm:"column:icon"`
	Sort             int    `json:"sort" gorm:"sort"`
	HideMenu         int    `json:"hideMenu" gorm:"hide_menu"`
	HideChildrenMenu int    `json:"hideChildrenMenu" gorm:"hide_children_menu"`
}

func (a Menu) TableName() string {
	return "sys_menu"
}
