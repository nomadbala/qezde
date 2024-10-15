package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qezde/api-gateway/internal/config"
	"github.com/qezde/api-gateway/pkg/server/response"
)

type ProxyHandler struct {
	config     config.Config
	httpClient *http.Client
}

func NewProxyHandler(config config.Config) *ProxyHandler {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &ProxyHandler{
		config:     config,
		httpClient: client,
	}
}

func (h *ProxyHandler) Routes(routerGroup fiber.Router, config config.Config) {
	routerGroup.All("/auth/*action", h.handleRequest(config.API.Auth))
}

func (h *ProxyHandler) handleRequest(targetURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Params("action")
		query := c.Request().URI().QueryString()
		target := targetURL + path
		if len(query) > 0 {
			target += "?" + string(query)
		}

		body := bytes.NewBuffer(c.Body())

		req, err := http.NewRequest(c.Method(), target, body)
		if err != nil {
			return response.InternalServerError(c, errors.New("failed to create request: "+err.Error()))
		}

		for key, values := range c.GetReqHeaders() {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		resp, err := h.httpClient.Do(req)
		if err != nil {
			return response.InternalServerError(c, errors.New("failed to forward request: "+err.Error()))
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)

		c.Set(fiber.HeaderContentType, resp.Header.Get(fiber.HeaderContentType))
		c.Status(resp.StatusCode)

		if _, err := io.Copy(c.Response().BodyWriter(), resp.Body); err != nil {
			return response.InternalServerError(c, errors.New("failed to copy response body: "+err.Error()))
		}

		return nil
	}
}
