package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	Resend  ResendConfig
	App     AppConfig
	Swagger SwaggerConfig
}

type ResendConfig struct {
	APIKey  string
	Sender  string
	Subject string
}

type AppConfig struct {
	Port string
}

type SwaggerConfig struct {
	BasePath string
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

	if err = envconfig.Process("RESEND", &config.Resend); err != nil {
		return
	}

	if err = envconfig.Process("APP", &config.App); err != nil {
		return
	}

	if err = envconfig.Process("SWAGGER", &config.Swagger); err != nil {
	}

	return
}
