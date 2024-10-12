package postgres

import (
	"context"
	"github.com/google/uuid"
	"qezde/user/db/postgres/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
	ctx     context.Context
}

func NewUserRepository(queries *sqlc.Queries, ctx context.Context) *UserRepository {
	return &UserRepository{queries: queries, ctx: ctx}
}

func (r UserRepository) GetAllUsers() (dest []sqlc.User, err error) {
	dest, err = r.queries.GetAllUsers(r.ctx)
	if err != nil {
		return
	}

	return
}

func (r UserRepository) GetUserById(id uuid.UUID) (dest sqlc.User, err error) {
	dest, err = r.queries.GetUserById(r.ctx, id)
	if err != nil {
		return
	}

	return
}

func (r UserRepository) CreateUser(params sqlc.CreateUserParams) (dest sqlc.User, err error) {
	dest, err = r.queries.CreateUser(r.ctx, params)
	if err != nil {
		return
	}

	return
}

func (r UserRepository) UpdateUser(params sqlc.UpdateUserParams) (dest sqlc.User, err error) {
	dest, err = r.queries.UpdateUser(r.ctx, params)
	if err != nil {
		return
	}

	return
}
