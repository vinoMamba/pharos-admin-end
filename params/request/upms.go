package request

type MenuSaveRequest struct {
	ParentId   string `json:"parentId"`
	MenuName   string `json:"menuName"`
	RoutePath  string `json:"routePath"`
	RouteName  string `json:"routeName"`
	Redirect   string `json:"redirect"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Type       int    `json:"type"`
	Icon       string `json:"icon"`
	Sort       int    `json:"sort"`
	Endpoint   int    `json:"endpoint"`
	Status     int    `json:"status"`
	HideMenu   int    `json:"hideMenu"`
}

type MenuUpdateRequest struct {
	ParentId   string `json:"parentId"`
	MenuName   string `json:"menuName"`
	RoutePath  string `json:"routePath"`
	RouteName  string `json:"routeName"`
	Redirect   string `json:"redirect"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Type       int    `json:"type"`
	Icon       string `json:"icon"`
	Sort       int    `json:"sort"`
	Endpoint   int    `json:"endpoint"`
	Status     int    `json:"status"`
	HideMenu   int    `json:"hideMenu"`
	MenuId     string `json:"menuId"`
}

type MenuDeleteRequest struct {
	MenuIds []string `json:"menuIds"`
}

type MenuDetailRequest struct {
	MenuId string `json:"menuId"`
}
