package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"qezde/notification/internal/domain/mail"
	"qezde/notification/pkg/errors"
)

func (h *Handler) Routes(r *mux.Router) {
	api := r.PathPrefix("/notification").Subrouter()
	{
		api.HandleFunc("/welcome_message", h.SendWelcomeMessage).Methods(http.MethodPost)
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	content := map[string]string{"status": "UP"}

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, "failed to encode healthcheck response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SendWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request mail.WelcomeMailRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "JSON_DECODING_ERROR: failed on decoding json from request", http.StatusBadRequest)
		return
	}

	if err := h.Resend.SendWelcomeEmail(request); err != errors.Nil {
		if err.Tag == errors.TagBadRequest {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if err.Tag == errors.TagInternalServerError {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
