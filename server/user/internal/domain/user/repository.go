package user

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetAllUsers(ctx context.Context) (dest []Entity, err error)
	GetUserById(ctx context.Context, id uuid.UUID) (dest *Entity, err error)
	CreateUser(ctx context.Context, request CreateUserRequest) (dest *Entity, err error)
	UpdateUser(ctx context.Context, id uuid.UUID, request UpdateUserRequest) (dest *Entity, err error)
}
