package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		APP        AppConfig
		API        APIConfig
		Middleware MiddlewareConfig
	}

	AppConfig struct {
		Port string
		Path string
	}

	APIConfig struct {
		Auth string
	}

	MiddlewareConfig struct {
		SigningKey string
	}
)

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

	signingKey := os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		return
	}

	config.Middleware = MiddlewareConfig{
		SigningKey: signingKey,
	}

	if err = envconfig.Process("API", &config.API); err != nil {
		return
	}

	return
}
