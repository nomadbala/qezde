package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"qezde/user/internal/domain/user"
)

func (h *Handler) Routes(r *mux.Router) {
	api := r.PathPrefix("/users").Subrouter()
	{
		api.HandleFunc("", h.GetAllUsers).Methods(http.MethodGet)
		api.HandleFunc("", h.CreateUser).Methods(http.MethodPost)
		api.HandleFunc("/{id}", h.GetUserById).Methods(http.MethodGet)
		api.HandleFunc("/{id}", h.UpdateUser).Methods(http.MethodPut)
	}
}

func (h *Handler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, map[string]string{"error": message})
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			h.respondWithError(w, http.StatusInternalServerError, "Failed to encode response")
		}
	}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	h.respondWithJSON(w, http.StatusOK, users)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.service.GetUserById(r.Context(), id)
	if err != nil {
		//if errors.Is(err, user.ErrNotFound) {
		//	h.respondWithError(w, http.StatusNotFound, "User not found")
		//	return
		//}
		h.respondWithError(w, http.StatusNotFound, "Failed to get user")
		return
	}

	h.respondWithJSON(w, http.StatusOK, user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.service.CreateUser(r.Context(), req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	h.respondWithJSON(w, http.StatusCreated, user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userID, err := uuid.Parse(id)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req user.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user, err := h.service.UpdateUser(r.Context(), userID, req)
	if err != nil {
		//if errors.Is(err, user.) {
		//	h.respondWithError(w, http.StatusNotFound, "User not found")
		//	return
		//}
		h.respondWithError(w, http.StatusNotFound, "Failed to update user")
		return
	}

	h.respondWithJSON(w, http.StatusOK, user)
}
