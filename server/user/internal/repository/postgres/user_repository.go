package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"qezde/user/internal/domain/user"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]user.Entity, error) {
	query := `
        SELECT id, username, password_hash, salt, email, email_verified, 
               first_name, last_name, date_of_birth, created_at, updated_at 
        FROM user_schema.users`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.Entity
	for rows.Next() {
		var u user.Entity
		err = rows.Scan(
			&u.ID, &u.Username, &u.PasswordHash, &u.Salt,
			&u.Email, &u.EmailVerified, &u.FirstName, &u.LastName,
			&u.DateOfBirth, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*user.Entity, error) {
	query := `
        SELECT id, username, password_hash, salt, email, email_verified, 
               first_name, last_name, date_of_birth, created_at, updated_at 
        FROM user_schema.users 
        WHERE id = $1 
        LIMIT 1`

	dest := &user.Entity{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&dest.ID, &dest.Username, &dest.PasswordHash, &dest.Salt,
		&dest.Email, &dest.EmailVerified, &dest.FirstName, &dest.LastName,
		&dest.DateOfBirth, &dest.CreatedAt, &dest.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}

	return dest, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, request user.CreateUserRequest) (*user.Entity, error) {
	query := `
        INSERT INTO user_schema.users (username, password_hash, salt, email) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id, username, password_hash, salt, email, email_verified, 
                  first_name, last_name, date_of_birth, created_at, updated_at`

	dest := &user.Entity{}
	err := r.pool.QueryRow(
		ctx,
		query,
		request.Username, request.PasswordHash, request.Salt, request.Email,
	).Scan(
		&dest.ID, &dest.Username, &dest.PasswordHash, &dest.Salt,
		&dest.Email, &dest.EmailVerified, &dest.FirstName, &dest.LastName,
		&dest.DateOfBirth, &dest.CreatedAt, &dest.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, request user.UpdateUserRequest) (*user.Entity, error) {
	query := `
        UPDATE user_schema.users 
        SET first_name = $1, last_name = $2, email = $3, 
            email_verified = $4, date_of_birth = $5 
        WHERE id = $6 
        RETURNING id, username, password_hash, salt, email, email_verified, 
                  first_name, last_name, date_of_birth, created_at, updated_at`

	dest := &user.Entity{}
	err := r.pool.QueryRow(
		ctx,
		query,
		request.FirstName, request.LastName, request.Email,
		request.EmailVerified, request.DateOfBirth, id,
	).Scan(
		&dest.ID, &dest.Username, &dest.PasswordHash, &dest.Salt,
		&dest.Email, &dest.EmailVerified, &dest.FirstName, &dest.LastName,
		&dest.DateOfBirth, &dest.CreatedAt, &dest.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}

	return dest, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM user_schema.users WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return user.ErrNotFound
	}

	return nil
}
