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

type MenuDeleteRequest struct {
	DeleteMenuIdList []string
}
