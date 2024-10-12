package handler

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"qezde/user/internal/config"
	"qezde/user/internal/handler/http"
	"qezde/user/internal/service"
	"qezde/user/pkg/server/response"
	"qezde/user/pkg/server/router"
	"time"
)

type Dependencies struct {
	Configs config.Config

	UserService service.UserService
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
			timeout.WithTimeout(60*time.Second),
			timeout.WithHandler(func(ctx *gin.Context) {
				ctx.Next()
			}),
			timeout.WithResponse(func(ctx *gin.Context) {
				response.StatusRequestTimeout(ctx)
			}),
		))

		h.HTTP.GET("/health", func(c *gin.Context) {
			c.JSON(200, "API Gateway is healthy :)")
		})

		userHandler := http.NewUserHandler(h.dependencies.UserService)

		api := h.HTTP.Group("/api/v1")
		{
			userHandler.Routes(api)
		}

		return
	}
}
