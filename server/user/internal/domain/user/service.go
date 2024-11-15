package user

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	GetAllUsers(ctx context.Context) (dest []DTO, err error)
	GetUserById(ctx context.Context, id string) (dest DTO, err error)
	CreateUser(ctx context.Context, request CreateUserRequest) (dest DTO, err error)
	UpdateUser(ctx context.Context, id uuid.UUID, request UpdateUserRequest) (dest DTO, err error)
}
