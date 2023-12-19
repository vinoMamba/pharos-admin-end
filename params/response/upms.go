package response

import "time"

type UserInfoResponse struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type RouterResponse struct {
	MenuId           string `json:"menuId"`
	MenuName         string `json:"menuName"`
	ParentId         string `json:"parentId"`
	RoutePath        string `json:"routePath"`
	RouteName        string `json:"routeName"`
	Redirect         string `json:"redirect"`
	Component        string `json:"component"`
	Type             int    `json:"type"`
	Affix            int    `json:"affix"`
	Icon             string `json:"icon"`
	Sort             int    `json:"sort"`
	HideChildrenMenu int    `json:"hideChildrenMenu"`
	HideMenu         int    `json:"hideMenu"`
}

type MenuListResponse struct {
	MenuId     string    `json:"menuId"`
	MenuName   string    `json:"menuName"`
	ParentId   string    `json:"parentId"`
	Icon       string    `json:"icon"`
	Permission string    `json:"permission"`
	Component  string    `json:"component"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
}

type MenuDetailResponse struct {
	Type       int    `json:"type"`
	MenuId     int64  `json:"menuId"`
	MenuName   string `json:"menuName"`
	ParentId   string `json:"parentId"`
	Icon       string `json:"icon"`
	Permission string `json:"permission"`
	Sort       int    `json:"sort"`
	RoutePath  string `json:"routePath"`
	RouteName  string `json:"routeName"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Affix      int    `json:"affix"`
	Status     int    `json:"status"`
	HideMenu   int    `json:"hideMenu"`
}

type RoleListResponse struct {
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
}

type RoleDetailResponse struct {
	RoleId   string   `json:"roleId"`
	RoleName string   `json:"roleName"`
	RoleCode string   `json:"roleCode"`
	Remark   string   `json:"remark"`
	Status   int      `json:"status"`
	MenuIds  []string `json:"menuIds"`
}
