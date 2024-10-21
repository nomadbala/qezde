package user

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	Username      string     `db:"username" json:"username"`
	PasswordHash  string     `db:"password_hash" json:"password_hash"`
	Salt          string     `db:"salt" json:"salt"`
	Email         string     `db:"email" json:"email"`
	EmailVerified *bool      `db:"email_verified" json:"email_verified"`
	FirstName     *string    `db:"first_name" json:"first_name"`
	LastName      *string    `db:"last_name" json:"last_name"`
	DateOfBirth   *time.Time `db:"date_of_birth" json:"date_of_birth"`
	CreatedAt     *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updated_at"`
}

func (e *Entity) ConvertToDTO() (dest DTO) {
	dest = DTO{
		ID:            e.ID,
		Username:      e.Username,
		Email:         e.Email,
		EmailVerified: e.EmailVerified,
		FirstName:     e.FirstName,
		LastName:      e.LastName,
		DateOfBirth:   e.DateOfBirth,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}

	return
}
