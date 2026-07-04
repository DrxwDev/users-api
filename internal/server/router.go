package server

import (
	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery(), gin.Logger())

	return router
}
