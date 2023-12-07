package handler

import "github.com/gin-gonic/gin"

func HandleUser(r *gin.Engine) {
	ug := r.Group("/user")
	ug.GET("/list", func(c *gin.Context) {})
}
