package users

import (
	"context"

	"github.com/google/uuid"

	"github.com/DrxwDev/users-api/internal/database"
	errormsg "github.com/DrxwDev/users-api/internal/errorMsg"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params CreateUserParams) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type postgres struct {
	db *database.Queries
}

func NewUserRepository(queries *database.Queries) UserRepository {
	return &postgres{
		db: queries,
	}
}

func (q postgres) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	user, err := q.db.Save(ctx, userParamsToDB(params))
	if err != nil {
		return User{}, errormsg.ErrUnableToSaveUser
	}

	return userRowFromDB(user), nil
}

func (q postgres) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	user, err := q.db.FindByID(ctx, id)
	if err != nil {
		return User{}, errormsg.ErrUserNotFound
	}

	return userFromDB(user), nil
}

func (q postgres) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := q.db.FindByEmail(ctx, email)
	if err != nil {
		return User{}, errormsg.ErrUserNotFound
	}

	return userFromDB(user), nil
}
