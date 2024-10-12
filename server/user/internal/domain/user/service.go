package user

import (
	"github.com/google/uuid"
	"qezde/user/db/postgres/sqlc"
)

type Service interface {
	GetAllUsers() (dest []DTO, err error)
	GetUserById(id uuid.UUID) (dest DTO, err error)
	CreateUser(params sqlc.CreateUserParams) (dest DTO, err error)
	UpdateUser(params sqlc.UpdateUserParams) (dest DTO, err error)
}
