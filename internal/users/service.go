package users

import (
	"context"

	errormsg "github.com/DrxwDev/users-api/internal/errorMsg"
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
	err := validateEmail(params.Email)
	if err != nil {
		return User{}, err
	}

	_, err = s.repo.GetUserByEmail(ctx, params.Email)
	if err == nil {
		return User{}, errormsg.ErrUserAlreadyExists
	}

	err = validatePassword(params.Password)
	if err != nil {
		return User{}, err
	}

	user, err := s.repo.CreateUser(ctx, params)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) GetUserByID(ctx context.Context, id string) (User, error) {
	err := validateUserID(id)
	if err != nil {
		return User{}, err
	}

	userID, err := parseUserID(id)
	if err != nil {
		return User{}, err
	}

	return s.repo.GetUserByID(ctx, userID)
}

func (s service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	err := validateEmail(email)
	if err != nil {
		return User{}, err
	}

	return s.repo.GetUserByEmail(ctx, email)
}
