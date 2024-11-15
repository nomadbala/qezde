package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	jwt.Claims
	UserId uuid.UUID `json:"user_id"`
}
