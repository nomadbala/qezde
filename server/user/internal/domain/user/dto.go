package user

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"qezde/user/db/postgres/sqlc"
)

type DTO struct {
	ID            uuid.UUID        `json:"id"`
	Username      string           `json:"username"`
	Email         string           `json:"email"`
	EmailVerified *bool            `json:"email_verified"`
	FirstName     *string          `json:"firstName"`
	LastName      *string          `json:"lastName"`
	DateOfBirth   pgtype.Date      `json:"date_of_birth"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

func ConvertEntityToDTO(user sqlc.User) (DTO, error) {
	dto := &DTO{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DateOfBirth:   user.DateOfBirth,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return dto, nil
}

func ConvertEntityToDTOs(users []sqlc.User) ([]DTO, error) {
	dtos := make([]DTO, len(users))

	for i, user := range users {
		dto, err := ConvertEntityToDTO(user)
		if err != nil {
			return nil, err
		}
		dtos[i] = *dto
	}

	return dtos, nil
}
