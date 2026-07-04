package users

import "github.com/DrxwDev/users-api/internal/auth"

type UserController struct {
	srv    UserService
	hasher auth.HashService
}

func NewUserController(srv UserService, hasher auth.HashService) UserController {
	return UserController{
		srv:    srv,
		hasher: hasher,
	}
}
