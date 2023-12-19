package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/vinoMamba.com/pharos-admin-end/logger"
	"github.com/vinoMamba.com/pharos-admin-end/middlewares"
	"github.com/vinoMamba.com/pharos-admin-end/models"
	"github.com/vinoMamba.com/pharos-admin-end/params/request"
	"github.com/vinoMamba.com/pharos-admin-end/params/response"
	"github.com/vinoMamba.com/pharos-admin-end/storage"
	"github.com/vinoMamba.com/pharos-admin-end/utils"
)

func HandleUpms(r *gin.Engine) {
	ug := r.Group("/upms")
	ug.Use(middlewares.AuthMiddleware).GET("/user/info", handleUserInfo)
	ug.Use(middlewares.AuthMiddleware).GET("/user/page_list", handleUserPageList)
	ug.Use(middlewares.AuthMiddleware).POST("/user/save", handleUserSave)
	ug.Use(middlewares.AuthMiddleware).PUT("/user/update", handleUserUpdate)
	ug.Use(middlewares.AuthMiddleware).DELETE("/user/delete", handleUserDelete)
	ug.Use(middlewares.AuthMiddleware).GET("/user/detail", handleUserDetail)

	ug.Use(middlewares.AuthMiddleware).GET("/menu/router", handleRouter)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/list", handleMenuList)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/detail", handleMenuDetail)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/tree", handleMenuList)
	ug.Use(middlewares.AuthMiddleware).GET("/getPermCode", handlePerCode)
	ug.Use(middlewares.AuthMiddleware).POST("/menu/save", handleMenuSave)
	ug.Use(middlewares.AuthMiddleware).PUT("/menu/update", handleMenuUpdate)
	ug.Use(middlewares.AuthMiddleware).DELETE("/menu/delete", handleMenuDelete)

	ug.Use(middlewares.AuthMiddleware).GET("/role/page_list", handleRolePageList)
	ug.Use(middlewares.AuthMiddleware).GET("/role/all_list", handleRoleAllList)
	ug.Use(middlewares.AuthMiddleware).POST("/role/save", handleRoleSave)
	ug.Use(middlewares.AuthMiddleware).PUT("/role/update", handleRoleUpdate)
	ug.Use(middlewares.AuthMiddleware).DELETE("/role/delete", handleRoleDelete)
	ug.Use(middlewares.AuthMiddleware).GET("/role/detail", handleRoleDetail)

}

func handleUserInfo(c *gin.Context) {
	log := logger.New(c)
	id := utils.GetCurrentUserId(c)
	u, err := storage.GetUserById(c, id)
	if err != nil {
		log.WithError(err).Errorln("没有该用户")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "没有该用户",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": response.UserInfoResponse{
			UserId:   u.UserId,
			Username: u.Username,
			RealName: u.RealName,
			Avatar:   u.Avatar,
			Password: u.Password,
		},
	})
}

func handleUserPageList(c *gin.Context) {
	log := logger.New(c)
	pageSize := c.Query("pageSize")
	pageNum := c.Query("pageNum")
	userList, total, err := storage.GetUserListByPage(c, cast.ToInt(pageSize), cast.ToInt(pageNum))
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}

	var items []response.UserListResponse
	for _, v := range userList {
		items = append(items, response.UserListResponse{
			UserId:      cast.ToString(v.UserId),
			Avatar:      v.Avatar,
			RealName:    v.RealName,
			Username:    v.Username,
			JobNumber:   v.JobNumber,
			Mobile:      v.Mobile,
			Email:       v.Email,
			HireDate:    v.HireDate,
			LeaveStatus: v.Status,
			LeaveTime:   cast.ToString(v.LeaveTime),
			AdminStatus: v.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"items":    items,
			"pageSize": pageSize,
			"pageNum":  pageNum,
			"total":    total,
		},
	})
}
func handleUserSave(c *gin.Context) {
	log := logger.New(c)
	var body request.UserCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	user := &models.User{
		UserId:    utils.GetSnowflakeIdInt64(),
		Username:  body.Username,
		Password:  body.Password,
		RealName:  body.RealName,
		JobNumber: body.JobNumber,
		Mobile:    body.Mobile,
		Email:     body.Email,
		Avatar:    body.Avatar,
		HireDate:  body.HireDate,
		Status:    body.LeaveStatus,
		LeaveTime: body.LeaveTime,
		IsAdmin:   body.AdminStatus,
	}

	if err := storage.CreateUser(c, user); err != nil {
		log.WithError(err).Errorln("添加失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "添加失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    nil,
	})
}
func handleUserUpdate(c *gin.Context) {

}
func handleUserDelete(c *gin.Context) {}
func handleUserDetail(c *gin.Context) {}

func handleRouter(c *gin.Context) {
	log := logger.New(c)
	var items []response.RouterResponse

	list, err := storage.GetRouteList(c)
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}

	for _, v := range list {
		items = append(items, response.RouterResponse{
			MenuId:           cast.ToString(v.MenuId),
			MenuName:         v.MenuName,
			ParentId:         v.ParentId,
			RoutePath:        v.RoutePath,
			RouteName:        v.RouteName,
			Redirect:         v.Redirect,
			Component:        v.Component,
			Type:             v.Type,
			Affix:            v.Affix,
			Icon:             v.Icon,
			Sort:             v.Sort,
			HideChildrenMenu: v.HideChildrenMenu,
			HideMenu:         v.HideMenu,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    items,
	})
}

func handleMenuList(c *gin.Context) {
	log := logger.New(c)
	var items []response.MenuListResponse
	list, err := storage.GetMenuList(c)
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}
	for _, v := range list {
		items = append(items, response.MenuListResponse{
			MenuId:     cast.ToString(v.MenuId),
			MenuName:   v.MenuName,
			ParentId:   v.ParentId,
			Icon:       v.Icon,
			Permission: v.Permission,
			Component:  v.Component,
			Sort:       v.Sort,
			Status:     v.Status,
			CreateTime: v.CreateTime,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    items,
	})
}

func handleMenuSave(c *gin.Context) {
	log := logger.New(c)
	var body request.MenuSaveRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	if err := storage.SaveMenu(c, &models.Menu{
		MenuId:     utils.GetSnowflakeIdInt64(),
		ParentId:   body.ParentId,
		MenuName:   body.MenuName,
		RoutePath:  body.RoutePath,
		RouteName:  body.RouteName,
		Redirect:   body.Redirect,
		Component:  body.Component,
		Permission: body.Permission,
		Type:       body.Type,
		Icon:       body.Icon,
		Sort:       body.Sort,
		Endpoint:   body.Endpoint,
		Status:     body.Status,
		HideMenu:   body.HideMenu,
	}); err != nil {
		log.WithError(err).Errorln("保存失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "保存失败",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	})
}

func handleMenuUpdate(c *gin.Context) {
	log := logger.New(c)
	var body request.MenuUpdateRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	if err := storage.UpdateMenu(c, &models.Menu{
		MenuId:     cast.ToInt64(body.MenuId),
		ParentId:   body.ParentId,
		MenuName:   body.MenuName,
		RoutePath:  body.RoutePath,
		RouteName:  body.RouteName,
		Redirect:   body.Redirect,
		Component:  body.Component,
		Permission: body.Permission,
		Type:       body.Type,
		Icon:       body.Icon,
		Sort:       body.Sort,
		Endpoint:   body.Endpoint,
		Status:     body.Status,
		HideMenu:   body.HideMenu,
	}); err != nil {
		log.WithError(err).Errorln("更新失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "更新失败",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	})
}

func handleMenuDelete(c *gin.Context) {
	log := logger.New(c)
	var body request.MenuDeleteRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	idList := utils.TransformStringToInt64(body.MenuIds)
	if err := storage.DeleteMenus(c, idList); err != nil {
		log.WithError(err).Errorln("删除失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "删除失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	})
}

func handleMenuDetail(c *gin.Context) {
	log := logger.New(c)
	menuId := c.Query("menuId")
	log.WithField("menuId", menuId).Infoln("menuId")

	m, err := storage.GetMenuById(c, cast.ToInt64(menuId))
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "没有该菜单",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": response.MenuDetailResponse{
			Type:       m.Type,
			MenuId:     m.MenuId,
			MenuName:   m.MenuName,
			ParentId:   m.ParentId,
			Icon:       m.Icon,
			Permission: m.Permission,
			Sort:       m.Sort,
			RoutePath:  m.RoutePath,
			RouteName:  m.RouteName,
			Component:  m.Component,
			Redirect:   m.Redirect,
			Affix:      m.Affix,
			Status:     m.Status,
			HideMenu:   m.HideMenu,
		},
	})
}

func handlePerCode(c *gin.Context) {
	log := logger.New(c)
	list, err := storage.GetPermCode(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}

	log.WithField("list", list).Infoln("list")
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    list,
	})
}

func handleRolePageList(c *gin.Context) {
	log := logger.New(c)
	pageSize := c.Query("pageSize")
	pageNum := c.Query("pageNum")
	log.WithField("pageSize", pageSize).WithField("pageNum", pageNum).Infoln("pageSize, pageNum")
	roleList, total, err := storage.GetRoleListByPage(c, cast.ToInt(pageNum), cast.ToInt(pageSize))
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}
	var items []response.RoleListResponse
	for _, v := range roleList {
		items = append(items, response.RoleListResponse{
			RoleId:   cast.ToString(v.RoleId),
			RoleName: v.RoleName,
			RoleCode: v.RoleCode,
			Remark:   v.Remark,
			Status:   v.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"items":    items,
			"pageSize": pageSize,
			"pageNum":  pageNum,
			"total":    total,
		},
	})
}

func handleRoleAllList(c *gin.Context) {
	log := logger.New(c)
	roleList, err := storage.GetAllRoleList(c)
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
		return
	}
	var list []response.RoleListResponse
	for _, v := range roleList {
		list = append(list, response.RoleListResponse{
			RoleId:   cast.ToString(v.RoleId),
			RoleName: v.RoleName,
			RoleCode: v.RoleCode,
			Remark:   v.Remark,
			Status:   v.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    list,
	})
}

func handleRoleSave(c *gin.Context) {
	log := logger.New(c)
	var body request.RoleCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	roleId := utils.GetSnowflakeIdInt64()
	role := &models.Role{
		RoleId:   roleId,
		RoleName: body.RoleName,
		RoleCode: body.RoleCode,
		Remark:   body.Remark,
		Status:   body.Status,
	}
	if err := storage.AddRole(c, role); err != nil {
		log.WithError(err).Errorln("添加失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "添加失败",
			"data":    nil,
		})
		return
	}
	if err := storage.CreateRoleMenus(c, roleId, body.MenuIds); err != nil {
		log.WithError(err).Errorln("添加失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "添加失败",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "添加成功",
		"data":    nil,
	})
}

func handleRoleUpdate(c *gin.Context) {
	log := logger.New(c)
	var body request.RoleUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	log.WithField("body", utils.Marshal(body)).Infoln("body")

	role, err := storage.GetRoleById(c, cast.ToInt64(body.RoleId))
	if err != nil {
		log.WithError(err).Errorln("没有该角色")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "没有该角色",
			"data":    nil,
		})
		return
	}
	role.RoleName = body.RoleName
	role.RoleCode = body.RoleCode
	role.Remark = body.Remark
	role.Status = body.Status

	log.WithField("role", utils.Marshal(role)).Infoln("body")

	if err := storage.UpdateRole(c, role); err != nil {
		log.WithError(err).Errorln("更新失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "更新失败",
			"data":    nil,
		})
		return
	}
	var roleIds []int64
	roleIds = append(roleIds, cast.ToInt64(body.RoleId))
	if err := storage.DeleteRoleMenuByRoleId(c, roleIds); err != nil {
		log.WithError(err).Errorln("更新失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "更新失败",
			"data":    nil,
		})
		return
	}

	if err := storage.CreateRoleMenus(c, cast.ToInt64(body.RoleId), body.MenuIds); err != nil {
		log.WithError(err).Errorln("更新失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "更新失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新成功",
		"data":    nil,
	})
}
func handleRoleDelete(c *gin.Context) {
	log := logger.New(c)
	var body request.RoleDeleteRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("参数错误")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	roleIds := utils.TransformStringToInt64(body.RoleIds)
	if err := storage.DeleteRole(c, roleIds); err != nil {
		log.WithError(err).Errorln("删除失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "删除失败",
			"data":    nil,
		})
		return
	}
	if err := storage.DeleteRoleMenuByRoleId(c, roleIds); err != nil {
		log.WithError(err).Errorln("删除失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "删除失败",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
		"data":    nil,
	})
}
func handleRoleDetail(c *gin.Context) {
	log := logger.New(c)
	roleId := c.Query("roleId")
	role, err := storage.GetRoleDetail(c, cast.ToInt64(roleId))
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "没有该角色",
			"data":    nil,
		})
		return
	}
	menuList, err := storage.GetMenuIdsByRoleId(c, cast.ToInt64(roleId))
	if err != nil {
		log.WithError(err).Errorln("查询失败")
		c.AbortWithStatusJSON(200, gin.H{
			"code":    1,
			"message": "没有该角色",
			"data":    nil,
		})
		return
	}
	var menuIds []string
	for _, v := range menuList {
		menuIds = append(menuIds, cast.ToString(v))
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": response.RoleDetailResponse{
			RoleId:   cast.ToString(role.RoleId),
			RoleName: role.RoleName,
			RoleCode: role.RoleCode,
			Remark:   role.Remark,
			Status:   role.Status,
			MenuIds:  menuIds,
		},
	})
}
