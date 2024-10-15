package handler

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"qezde/auth/internal/config"
	"qezde/auth/internal/handler/http"
	"qezde/auth/internal/service"
	"qezde/auth/pkg/server/response"
	"qezde/auth/pkg/server/router"
)

type Dependencies struct {
	Configs config.Config

	AuthenticationService *service.AuthenticationService
}

type Handler struct {
	dependencies Dependencies
	HTTP         *gin.Engine
}

type Configuration func(h *Handler) error

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
	}

	for _, cfg := range configs {
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = router.New()

		h.HTTP.Use(timeout.New(
			timeout.WithTimeout(h.dependencies.Configs.App.Timeout),
			timeout.WithHandler(func(c *gin.Context) {
				c.Next()
			}),
			timeout.WithResponse(func(c *gin.Context) {
				response.StatusRequestTimeout(c)
			}),
		))

		authHandler := http.NewAuthenticationHandler(h.dependencies.AuthenticationService)

		api := h.HTTP.Group("")
		{
			authHandler.Routes(api)
		}

		return
	}
}
