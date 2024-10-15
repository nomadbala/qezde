package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Object struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

func UnAuthorized(c *fiber.Ctx, err error) error {
	h := Object{
		Success: false,
		Message: err.Error(),
	}

	return c.Status(http.StatusUnauthorized).JSON(h)
}

func InternalServerError(c *fiber.Ctx, err error) error {
	h := Object{
		Success: false,
		Message: err.Error(),
	}
	return c.Status(http.StatusInternalServerError).JSON(h)
}

func MethodNotAllowedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		allowedMethods := map[string]bool{
			http.MethodGet:    true,
			http.MethodPost:   true,
			http.MethodPut:    true,
			http.MethodDelete: true,
		}

		if !allowedMethods[c.Method()] {
			return c.Status(http.StatusMethodNotAllowed).JSON(fiber.Map{"error": "Method Not Allowed"})
		}

		return c.Next()
	}
}
