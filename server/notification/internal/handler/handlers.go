package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"qezde/notification/internal/domain/mail"
)

func (h *Handler) Routes(r *mux.Router) {
	api := r.PathPrefix("/notification").Subrouter()
	{
		api.HandleFunc("/welcome_message", h.SendWelcomeMessage).Methods(http.MethodPost)
	}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Check the health of the service
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Healthy"}
	_ = json.NewEncoder(w).Encode(response)
}

// SendWelcomeMessage godoc
// @Summary Send Welcome Message
// @Description Sends a welcome email to the user with the provided email and code.
// @Tags notifications
// @Accept json
// @Produce json
// @Param request body mail.WelcomeMailRequest true "Welcome Mail Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /notification/welcome_message [post]
func (h *Handler) SendWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request mail.WelcomeMailRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.Resend.SendWelcomeEmail(request.Email, request.Code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Welcome message sent successfully"}
	_ = json.NewEncoder(w).Encode(response)
}
