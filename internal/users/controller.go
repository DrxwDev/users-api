package users

type UserController struct {
	srv UserService
}

func NewUserController(srv UserService) UserController {
	return UserController{
		srv: srv,
	}
}
