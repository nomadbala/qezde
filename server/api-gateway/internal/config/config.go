package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APP        AppConfig
	API        APIConfig
	CORS       CORSConfig
	Middleware MiddlewareConfig
	Keycloak   KeycloakConfig
}

type AppConfig struct {
	Port string
	Path string
}

type APIConfig struct {
	Auth         string
	Notification string
}

type MiddlewareConfig struct {
	SigningKey string
}

type CORSConfig struct {
	AllowedOrigins []string
}

type KeycloakConfig struct {
	Host         string
	ClientId     string
	ClientSecret string
	Realm        string
}

func New() (config Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return
	}

	if err = envconfig.Process("APP", &config.APP); err != nil {
		return
	}

	if err = envconfig.Process("MIDDLEWARE", &config.Middleware); err != nil {
		return
	}

	if err = envconfig.Process("CORS", &config.CORS); err != nil {
		return
	}

	if err = envconfig.Process("API", &config.API); err != nil {
		return
	}

	if err = envconfig.Process("KEYCLOAK", &config.Keycloak); err != nil {
		return
	}

	return
}
