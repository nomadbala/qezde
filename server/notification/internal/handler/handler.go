package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
	"qezde/notification/internal/service"
	"qezde/notification/pkg/errors"
)

type Configuration func(h *Handler) errors.Error

type Handler struct {
	Router *mux.Router

	Resend *service.Service
}

func New(configs ...Configuration) (h *Handler, err errors.Error) {
	h = &Handler{}

	for _, cfg := range configs {
		if err = cfg(h); err != errors.Nil {
			fmt.Println(cfg(h))
			return h, errors.New("HANDLER_ERROR", fmt.Sprintf("%s", cfg(h)))
		}
	}

	return h, errors.Nil
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err errors.Error) {
		h.Router = mux.NewRouter()

		h.Router.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

		h.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		api := h.Router.PathPrefix("/api/v1").Subrouter()

		h.Routes(api)

		if h.Router == nil {
			return errors.New("HANDLER_ERROR", "failed while initializing mux router")
		}

		return errors.Nil
	}
}

func WithResendService(resend *service.Service) Configuration {
	return func(h *Handler) (err errors.Error) {
		h.Resend = resend

		if h.Resend == nil {
			return errors.New("HANDLER_ERROR", "failed to initialize resend service in handler")
		}

		return
	}
}
