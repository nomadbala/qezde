package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"

	"qezde/api-gateway/internal/domain"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"qezde/api-gateway/internal/config"
	"qezde/api-gateway/pkg/log"
)

type Dependencies struct {
	Configs config.Config
}

type Handler struct {
	dependencies Dependencies
	router       *mux.Router
	httpClient   *http.Client
}

func New(d Dependencies) *mux.Router {
	h := &Handler{
		dependencies: d,
		router:       mux.NewRouter(),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	h.router.Use(h.corsMiddleware)
	h.router.Use(h.loggingMiddleware)

	h.router.HandleFunc("/health", h.healthCheck).Methods(http.MethodGet)

	authRouter := h.router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/{action}", h.proxyHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch)

	api := h.router.PathPrefix("/api/v1").Subrouter()

	api.Use(h.authMiddleware)

	notificationRouter := api.PathPrefix("/notification").Subrouter()
	notificationRouter.HandleFunc("/{action}", h.proxyHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch)

	return h.router
}

func (h *Handler) proxyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	action := mux.Vars(r)["action"]
	servicePath := strings.TrimPrefix(r.URL.Path, "/notification/")

	var targetURL string
	if strings.Contains(r.URL.Path, "/notification/") {
		targetURL = h.dependencies.Configs.API.Notification
	} else if strings.Contains(r.URL.Path, "/auth/") {
		targetURL = h.dependencies.Configs.API.Auth
	} else {
		http.Error(w, "Unknown service", http.StatusBadRequest)
		return
	}

	fullURL, err := url.Parse(targetURL + "/" + action)
	if err != nil {
		logger.Log.Errorw("URL parse error",
			"error", err,
			"target", targetURL,
			"action", action,
		)
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		return
	}

	fullURL.RawQuery = r.URL.RawQuery

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.Errorw("Request body read error", "error", err)
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest(r.Method, fullURL.String(), bytes.NewBuffer(body))
	if err != nil {
		logger.Log.Errorw("Request creation error",
			"error", err,
			"method", r.Method,
			"url", fullURL.String(),
		)
		http.Error(w, "Unable to create request", http.StatusInternalServerError)
		return
	}

	safeCopyHeaders(r.Header, req.Header)

	logger.Log.Infow("Proxying request",
		"method", r.Method,
		"url", fullURL.String(),
		"service_path", servicePath,
	)

	resp, err := h.httpClient.Do(req)
	if err != nil {
		logger.Log.Errorw("Backend service request failed",
			"error", err,
			"url", fullURL.String(),
		)
		http.Error(w, "Request to backend service failed: ", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		logger.Log.Errorw("Response body copy error", "error", err)
	}
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := "Healthy"
	_ = json.NewEncoder(w).Encode(response)
}

func safeCopyHeaders(src http.Header, dst http.Header) {
	safeHeaders := []string{
		"Accept",
		"Content-Type",
		"Authorization",
		"User-Agent",
	}

	for _, header := range safeHeaders {
		if values := src.Values(header); len(values) > 0 {
			dst[header] = values
		}
	}
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
		return uuid.Nil, errors.New("failed to parse token: ")
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
