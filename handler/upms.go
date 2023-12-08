package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba.com/pharos-admin-end/logger"
	"github.com/vinoMamba.com/pharos-admin-end/middlewares"
	"github.com/vinoMamba.com/pharos-admin-end/params/response"
	"github.com/vinoMamba.com/pharos-admin-end/storage"
	"github.com/vinoMamba.com/pharos-admin-end/utils"
)

func HandleUpms(r *gin.Engine) {
	ug := r.Group("/upms")
	ug.Use(middlewares.AuthMiddleware).GET("/user/info", handleUserInfo)
	ug.Use(middlewares.AuthMiddleware).GET("/menu/router", handleRouter)
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
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    list,
	})
}
