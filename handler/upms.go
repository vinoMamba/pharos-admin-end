package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	ug.Use(middlewares.AuthMiddleware).GET("/menu/router", handleRouter)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/list", handleMenuList)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/save", handleMenuSave)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/delete", handleMenuDelete)
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
			MenuId:           v.MenuId,
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
			MenuId:     v.MenuId,
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

	idList := utils.TransformStringToInt64(body.DeleteMenuIdList)
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
