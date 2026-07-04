package users

import (
	"github.com/google/uuid"

	"github.com/DrxwDev/users-api/internal/database"
)

func userFromDB(user database.User) User {
	return User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func userRowFromDB(user database.SaveRow) User {
	return User{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func userParamsToDB(params CreateUserParams) database.SaveParams {
	return database.SaveParams{
		ID:       uuid.New(),
		Email:    params.Email,
		Password: params.Password,
	}
}

func userParamsToDomain(params UserRequest) CreateUserParams {
	return CreateUserParams(params)
}

func userDomainToDTO(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.Password,
	}
}
