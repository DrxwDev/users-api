package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv UserService
}

func NewUserController(srv UserService) UserController {
	return UserController{
		srv: srv,
	}
}

func (c UserController) CreateUser(ctx *gin.Context) {
	var payload UserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid user payload",
			"error":   err.Error(),
		})
		return
	}

	params := userRequestToDomain(payload)
	user, err := c.srv.CreateUser(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  "failed",
			"message": "conflic while creating user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "user created successfully",
		"user":    userDomainToDTO(user),
	})
}
