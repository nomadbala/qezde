package handler

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"net/http"
	logger "qezde/api-gateway/pkg/log"
	"strings"
	"time"
)

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := h.dependencies.Configs.CORS.AllowedOrigins
		origin := r.Header.Get("Origin")

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		var originAllowed bool
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				originAllowed = true
				break
			}
		}

		if !originAllowed && origin != "" {
			logger.Log.Warnf("Blocked CORS from origin: %s", origin)
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Log.Infow("Incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		logger.Log.Infow("Request processed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration,
		)
	})
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(h.dependencies.Configs.Keycloak.Host)
		fmt.Println(h.dependencies.Configs.Keycloak.Realm)
		fmt.Println(h.dependencies.Configs.Keycloak.ClientId)
		fmt.Println(h.dependencies.Configs.Keycloak.ClientSecret)

		client := gocloak.NewClient(h.dependencies.Configs.Keycloak.Host)
		ctx := context.Background()

		header := r.Header.Get("Authorization")
		if err := h.validateAuthorizationHeader(header); err != nil {
			http.Error(w, "No valid authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		rptResult, err := client.RetrospectToken(ctx, token, h.dependencies.Configs.Keycloak.ClientId, h.dependencies.Configs.Keycloak.ClientSecret, h.dependencies.Configs.Keycloak.Realm)
		if err != nil {
			http.Error(w, "Un authorized", http.StatusUnauthorized)
			return
		}

		if !*rptResult.Active {
			http.Error(w, "Un authorized", http.StatusUnauthorized)
		}

		if *rptResult.Active {
			next.ServeHTTP(w, r)
		}
	})
}

//
// =========================================
//
//	func (h *Handler) authMiddleware(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			client := gocloak.NewClient(h.dependencies.Configs.Keycloak.Host)
//			ctx := context.Background()
//
//			// Получаем токен из заголовка запроса
//			header := r.Header.Get("Authorization")
//			if err := h.validateAuthorizationHeader(header); err != nil {
//				http.Error(w, err.Error(), http.StatusUnauthorized)
//				return
//			}
//
//			userToken := strings.TrimPrefix(header, "Bearer ")
//
//			// Используем учетные данные клиента для проверки пользовательского токена
//			rptResult, err := client.RetrospectToken(
//				ctx,
//				userToken,
//				"qezde",
//				"1AXaXkNFE6PPeLkt7xGPmvNxNIdcle0G",
//				"qezde",
//			)
//			if err != nil {
//				log.Printf("Error introspecting token: %v", err)
//				log.Println(userToken)
//				log.Println(h.dependencies.Configs.Keycloak.ClientId)
//				log.Println(h.dependencies.Configs.Keycloak.ClientSecret)
//				log.Println(h.dependencies.Configs.Keycloak.Realm)
//				http.Error(w, "Token validation failed", http.StatusUnauthorized)
//				return
//			}
//
//			if !*rptResult.Active {
//				http.Error(w, "Token is not active", http.StatusUnauthorized)
//				return
//			}
//
//			// Добавляем информацию о пользователе в контекст запроса
//			userInfo, err := client.GetUserInfo(
//				ctx,
//				userToken,
//				h.dependencies.Configs.Keycloak.Realm,
//			)
//			if err != nil {
//				log.Printf("Error getting user info: %v", err)
//				http.Error(w, "Failed to get user info", http.StatusUnauthorized)
//				return
//			}
//
//			// Создаем обогащенный контекст с информацией о пользTователе
//			enrichedContext := context.WithValue(r.Context(), "user_info", userInfo)
//			next.ServeHTTP(w, r.WithContext(enrichedContext))
//		})
//	}
//
