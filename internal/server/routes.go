package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DrxwDev/users-api/internal/users"
)

func Routes(router *gin.Engine, controller users.UserController) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Hello World",
		})
	})

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/users", controller.CreateUser)
	apiV1.POST("/users", controller.GetUserByEmail)
	apiV1.GET("/users/:id", controller.GetUserByID)
}
