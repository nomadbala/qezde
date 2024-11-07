package handler

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
	"qezde/notification/internal/service"
)

type Configuration func(h *Handler) error

type Handler struct {
	Router *mux.Router

	Resend *service.Service
}

func New(configs ...Configuration) (h *Handler, err error) {
	h = &Handler{}

	for _, cfg := range configs {
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.Router = mux.NewRouter()

		h.Router.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

		h.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		api := h.Router.PathPrefix("/api/v1").Subrouter()

		h.Routes(api)

		return
	}
}

func WithResendService(resend *service.Service) Configuration {
	return func(h *Handler) (err error) {
		h.Resend = resend

		return
	}
}
