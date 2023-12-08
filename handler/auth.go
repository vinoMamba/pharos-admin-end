package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba.com/pharos-admin-end/config"
	"github.com/vinoMamba.com/pharos-admin-end/logger"
	"github.com/vinoMamba.com/pharos-admin-end/middlewares"
	"github.com/vinoMamba.com/pharos-admin-end/params/request"
	"github.com/vinoMamba.com/pharos-admin-end/params/response"
	"github.com/vinoMamba.com/pharos-admin-end/storage"
	"github.com/vinoMamba.com/pharos-admin-end/utils"
)

func HandleUser(r *gin.Engine) {
	ug := r.Group("/auth")
	ug.POST("/login/dingtalk", handleDingtalkLogin)
	ug.POST("/login/password", handlePwdLogin)
	ug.Use(middlewares.AuthMiddleware).POST("/login/refresh", handleRefresh)

}

func handleDingtalkLogin(c *gin.Context) {
	log := logger.New(c)
	var body request.DingtalkLoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json error")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
}

func handlePwdLogin(c *gin.Context) {
	log := logger.New(c)
	var body request.PwdLoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json error")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	u, err := storage.GetUserByUsername(c, body.Username)
	if err != nil {
		log.WithError(err).Errorln("没有该用户")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "用户名不正确",
			"data":    nil,
		})
		return
	}
	if ok := utils.CheckHashPassword(u.Password, body.Password); !ok {
		log.WithError(err).Errorln("密码错误")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "密码错误",
			"data":    nil,
		})
		return
	}

	token, err := utils.CreateJwt(u.UserId, u.Username)
	if err != nil {
		log.WithError(err).Errorln("密码错误")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "服务器错误",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": response.LoginResponse{
			AccessToken: token,
			ExpiresIn:   int(config.GetJwtExpiresIn()),
		},
	})
}

func handleRefresh(c *gin.Context) {
	log := logger.New(c)

	id := utils.GetCurrentUserId(c)

	u, err := storage.GetUserById(c, id)
	if err != nil {
		log.WithError(err).Errorln("没有该用户")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "服务器错误",
			"data":    nil,
		})
		return
	}

	token, err := utils.CreateJwt(u.UserId, u.Username)
	if err != nil {
		log.WithError(err).Errorln("创建token失败")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "服务器错误",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": response.LoginResponse{
			AccessToken: token,
			ExpiresIn:   int(config.GetJwtExpiresIn()),
		},
	})
}
