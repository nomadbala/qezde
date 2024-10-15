package http

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/qezde/api-gateway/internal/domain"
	"github.com/qezde/api-gateway/pkg/server/response"
)

func (h *ProxyHandler) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if err := h.validateAuthorizationHeader(header); err != nil {
			return response.UnAuthorized(c, err)
		}

		token := strings.TrimPrefix(header, "Bearer ")
		userId, err := h.ParseToken(token)
		if err != nil {
			return response.UnAuthorized(c, err)
		}

		c.Locals("userId", userId)
		return c.Next()
	}
}

func (h *ProxyHandler) validateAuthorizationHeader(header string) error {
	if header == "" {
		return errors.New("missing authorization header")
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return errors.New("invalid authorization header format")
	}
	return nil
}

func (h *ProxyHandler) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &domain.TokenClaims{}, h.getTokenSigningKey())
	if err != nil {
		return uuid.Nil, errors.New("failed to parse token: " + err.Error())
	}

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return uuid.Nil, errors.New("invalid token")
}

func (h *ProxyHandler) getTokenSigningKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(h.config.Middleware.SigningKey), nil
	}
}
