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
	MenuId           int64  `json:"menuId"`
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
