package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/qezde/api-gateway/internal/config"
	"github.com/qezde/api-gateway/internal/handler/http"
	"github.com/qezde/api-gateway/pkg/server/router"
)

type Dependencies struct {
	Configs config.Config
}

type Handler struct {
	dependencies Dependencies
	HTTP         *fiber.App
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

		h.HTTP.Use(healthcheck.New(healthcheck.Config{
			LivenessProbe: func(c *fiber.Ctx) bool {
				return true
			},
			LivenessEndpoint: "/live",
			ReadinessProbe: func(c *fiber.Ctx) bool {
				return true
			},
			ReadinessEndpoint: "/ready",
		}))

		h.HTTP.Use(csrf.New())

		proxy := http.NewProxyHandler(h.dependencies.Configs)

		api := h.HTTP.Group(h.dependencies.Configs.APP.Path)
		{
			proxy.Routes(api, h.dependencies.Configs)
		}

		return
	}
}
