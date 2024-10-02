package handler

import (
	"auth/internal/config"
	"auth/internal/handler/http"
	"auth/internal/service"
	"auth/pkg/server/response"
	"auth/pkg/server/router"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

		h.HTTP.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		h.HTTP.GET("/health", func(c *gin.Context) {
			response.OK(c, gin.H{})
		})

		authHandler := http.NewAuthenticationHandler(h.dependencies.AuthenticationService)

		api := h.HTTP.Group(h.dependencies.Configs.App.Path)
		{
			authHandler.Routes(api)
		}

		return
	}
}
