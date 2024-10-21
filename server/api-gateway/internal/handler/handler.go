package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"qezde/api-gateway/internal/config"
	"qezde/api-gateway/internal/domain"
)

type Dependencies struct {
	Configs config.Config
}

type Handler struct {
	dependencies Dependencies
	router       *mux.Router
}

func New(d Dependencies) *mux.Router {
	h := &Handler{
		dependencies: d,
		router:       mux.NewRouter(),
	}

	h.router.HandleFunc("/health", h.healthCheckHandler).Methods("GET")
	h.router.HandleFunc("/auth/{action}", h.proxyHandler).Methods("GET")
	//h.router.Use(h.middleware)

	return h.router
}

func (h *Handler) proxyHandler(w http.ResponseWriter, r *http.Request) {
	action := mux.Vars(r)["action"]
	targetURL := h.dependencies.Configs.API.Auth + "/" + action

	method := r.Method
	query := r.URL.RawQuery

	if query != "" {
		targetURL += "?" + query
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest(method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "unable to create request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header.Clone()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "request to backend service failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	io.Copy(w, resp.Body)
}

func (h *Handler) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "UP"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if err := h.validateAuthorizationHeader(header); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		userId, err := h.parseToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) validateAuthorizationHeader(header string) error {
	if header == "" {
		return errors.New("missing authorization header")
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return errors.New("invalid authorization header format")
	}
	return nil
}

func (h *Handler) parseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &domain.TokenClaims{}, h.getTokenSigningKey())
	if err != nil {
		return uuid.Nil, errors.New("failed to parse token: " + err.Error())
	}

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return uuid.Nil, errors.New("invalid token")
}

func (h *Handler) getTokenSigningKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(h.dependencies.Configs.Middleware.SigningKey), nil
	}
}
