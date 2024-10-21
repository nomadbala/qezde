package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"qezde/user/internal/domain/user"
)

type UserRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewUserRepository(conn *pgx.Conn, ctx context.Context) *UserRepository {
	return &UserRepository{conn: conn, ctx: ctx}
}

func (r UserRepository) GetAllUsers() (dest []user.Entity, err error) {
	rows, err := r.conn.Query(r.ctx, "SELECT id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at FROM user_schema.users")
	if err != nil {
		return dest, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user.Entity
		if err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Salt, &user.Email, &user.EmailVerified, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return dest, err
		}

		dest = append(dest, user)
	}

	return
}

func (r UserRepository) GetUserById(id uuid.UUID) (dest user.Entity, err error) {
	err = r.conn.QueryRow(
		r.ctx,
		"SELECT id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at FROM user_schema.users WHERE id=$1 LIMIT 1", id).
		Scan(&dest.ID, &dest.Username, &dest.PasswordHash, &dest.Salt, &dest.Email, &dest.EmailVerified, &dest.FirstName, &dest.LastName, &dest.DateOfBirth, &dest.CreatedAt, &dest.UpdatedAt)
	if err != nil {
		return
	}

	return
}

func (r UserRepository) CreateUser(request user.CreateUserRequest) (dest user.Entity, err error) {
	err = r.conn.QueryRow(
		r.ctx,
		"INSERT INTO user_schema.users (username, password_hash, salt, email) VALUES ($1, $2, $3, $4)",
		request.Username, request.PasswordHash, request.Salt, request.Email,
	).Scan(&dest.ID, &dest.Username, &dest.PasswordHash, &dest.Salt, &dest.Email, &dest.EmailVerified, &dest.FirstName, &dest.LastName, &dest.DateOfBirth, &dest.CreatedAt, &dest.UpdatedAt)
	if err != nil {
		return
	}

	return
}

func (r UserRepository) UpdateUser(request user.CreateUserRequest) (dest user.Entity, err error) {
	err = r.conn.QueryRow(
		r.ctx,
		"UPDATE user_schema.users SET first_name = $1, last_name = $2, email = $3, email_verified = $4, date_of_birth = $5 WHERE id=$6 RETURNING id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at",
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return
	}

	return
}
