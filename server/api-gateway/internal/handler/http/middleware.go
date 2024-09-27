package http

import (
	"errors"
	"gateway/internal/domain"
	"gateway/pkg/server/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func (h *ProxyHandler) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if err := h.validateAuthorizationHeader(header); err != nil {
			response.UnAuthorized(c, err)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		userId, err := h.ParseToken(token)
		if err != nil {
			response.UnAuthorized(c, err)
			return
		}

		c.Set("userId", userId)
	}
}

func (h *ProxyHandler) validateAuthorizationHeader(header string) error {
	if header == "" {
		return errors.New("authorization header is missing")
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return errors.New("authorization header is invalid")
	}
	return nil
}

func (h *ProxyHandler) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &domain.TokenClaims{}, h.getTokenSigningKey())
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	return claims.UserId, nil
}

func (h *ProxyHandler) getTokenSigningKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(h.config.Middleware.SigningKey), nil
	}
}
