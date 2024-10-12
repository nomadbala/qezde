package user

import (
	"github.com/google/uuid"
	"qezde/user/db/postgres/sqlc"
)

type Repository interface {
	GetAllUsers() (dest []sqlc.User, err error)
	GetUserById(id uuid.UUID) (dest sqlc.User, err error)
	CreateUser(params sqlc.CreateUserParams) (dest sqlc.User, err error)
	UpdateUser(params sqlc.UpdateUserParams) (dest sqlc.User, err error)
}
