package users

import (
	"context"

	"github.com/google/uuid"

	"github.com/DrxwDev/users-api/internal/database"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params CreateUserParams) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type postgres struct {
	q *database.Queries
}

func NewUserRepository(queries *database.Queries) UserRepository {
	return &postgres{
		q: queries,
	}
}

func (db postgres) CreateUser(ctx context.Context, params CreateUserParams) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (db postgres) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (db postgres) GetUserByEmail(ctx context.Context, email string) (User, error) {
	panic("not implemented") // TODO: Implement
}
