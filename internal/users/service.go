package users

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, params CreateUserParams) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &service{
		repo: repo,
	}
}

func (s service) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (s service) GetUserByID(ctx context.Context, id string) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (s service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	panic("not implemented") // TODO: Implement
}
