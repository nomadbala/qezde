package domain

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

var (
	UsernameIsRequiredException = errors.New("username in request is required")
	PasswordIsRequiredException = errors.New("password in request is required")
	EmailIsRequiredException    = errors.New("email in request is required")
)

type RegistrationRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (r *RegistrationRequest) Validate() error {
	if r.Username == "" {
		return UsernameIsRequiredException
	}

	if r.Password == "" {
		return PasswordIsRequiredException
	}

	if r.Email == "" {
		return EmailIsRequiredException
	}

	return nil
}

func (r *RegistrationRequest) ToJSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r *UserDTO) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), r)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginRequest) Validate() error {
	if l.Username == "" {
		return UsernameIsRequiredException
	}

	if l.Password == "" {
		return PasswordIsRequiredException
	}

	return nil
}

type CreateUserRequest struct {
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
	Salt         string `json:"salt" binding:"required"`
	Email        string `json:"email" binding:"required"`
}
