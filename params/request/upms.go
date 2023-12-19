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

type RoleCreateRequest struct {
	RoleName string   `json:"roleName"`
	RoleCode string   `json:"roleCode"`
	Remark   string   `json:"remark"`
	Status   int      `json:"status"`
	MenuIds  []string `json:"menuIds"`
}

type RoleUpdateRequest struct {
	RoleId   string   `json:"roleId"`
	RoleName string   `json:"roleName"`
	RoleCode string   `json:"roleCode"`
	Remark   string   `json:"remark"`
	Status   int      `json:"status"`
	MenuIds  []string `json:"menuIds"`
}

type RoleDeleteRequest struct {
	RoleIds []string `json:"roleIds"`
}
