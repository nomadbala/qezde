package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"qezde/user/internal/config"
	"qezde/user/internal/domain/user"
)

type Dependencies struct {
	Configs     config.Config
	UserService user.Service
	Ctx         context.Context
}

type Handler struct {
	dependencies Dependencies
	Router       *mux.Router
	service      user.Service
	Ctx          context.Context
}

type Configuration func(h *Handler) error

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
		Router:       mux.NewRouter(),
		service:      d.UserService,
		Ctx:          d.Ctx,
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
		h.Router.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

		h.Router.Use(h.corsMiddleware)

		api := h.Router.PathPrefix("/api/v1").Subrouter()
		h.Routes(api)

		return
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode("Healthy")
}
