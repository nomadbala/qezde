package user

import (
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"time"
)

type DTO struct {
	ID            uuid.UUID  `json:"id"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	EmailVerified *bool      `json:"email_verified"`
	FirstName     *string    `json:"first_name"`
	LastName      *string    `json:"last_name"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PasswordHash string `db:"password_hash" json:"password_hash" binding:"required"`
	Salt         string `db:"salt" json:"salt" binding:"required"`
}

const (
	CreateUserRequestUsernameMinLength = 3
	CreateUserRequestUsernameMaxLength = 32
)

func (r *CreateUserRequest) Validate() (err error) {
	if len(r.Username) < CreateUserRequestUsernameMinLength {
		return fmt.Errorf("username must be longer than 3 characters")
	}

	if len(r.Username) > CreateUserRequestUsernameMaxLength {
		return fmt.Errorf("username must not exceed 32 characters")
	}

	_, err = mail.ParseAddress(r.Email)
	if err != nil {
		return fmt.Errorf("invalid email")
	}

	if (len(r.PasswordHash)) == 0 || len(r.Salt) == 0 {
		return fmt.Errorf("invalid password hash or salt")
	}

	return
}

type UpdateUserRequest struct {
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	EmailVerified *bool      `json:"email_verified"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
}
