package server

import "github.com/gin-gonic/gin"

func SetupServer() *gin.Engine {
	r := gin.Default()
	return r
}
