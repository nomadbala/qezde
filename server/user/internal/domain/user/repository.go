package user

import (
	"github.com/google/uuid"
)

type Repository interface {
	GetAllUsers() (dest []Entity, err error)
	GetUserById(id uuid.UUID) (dest Entity, err error)
	CreateUser(request CreateUserRequest) (dest Entity, err error)
	UpdateUser(request CreateUserRequest) (dest Entity, err error)
}
