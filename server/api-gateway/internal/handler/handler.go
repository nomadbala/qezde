package handler

import (
	"gateway/internal/config"
	"gateway/internal/handler/http"
	"gateway/pkg/server/response"
	"gateway/pkg/server/router"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"time"
)

type Dependencies struct {
	Configs config.Config
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

		proxy := http.NewProxyHandler(h.dependencies.Configs)

		api := h.HTTP.Group("/auth")
		{
			proxy.Routes(api, h.dependencies.Configs)
		}

		return
	}
}
