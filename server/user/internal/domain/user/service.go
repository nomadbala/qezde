package user

import (
	"github.com/google/uuid"
)

type Service interface {
	GetAllUsers() (dest []DTO, err error)
	GetUserById(id uuid.UUID) (dest DTO, err error)
	CreateUser(request CreateUserRequest) (dest DTO, err error)
	UpdateUser(request CreateUserRequest) (dest DTO, err error)
}
