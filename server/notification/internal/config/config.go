package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
	"qezde/notification/pkg/errors"
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

func New() (Config, errors.Error) {
	config := Config{}

	root, err := os.Getwd()
	if err != nil {
		return config, errors.New("CONFIG_ERROR", "failed to get working directory", errors.TagInternalServerError)
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return config, errors.New("CONFIG_ERROR", "failed to load .env file", errors.TagInternalServerError)
	}

	if err = envconfig.Process("RESEND", &config.Resend); err != nil {
		return config, errors.New("CONFIG_ERROR", "failed to load .env resend variables", errors.TagInternalServerError)
	}

	if err = envconfig.Process("APP", &config.App); err != nil {
		return config, errors.New("CONFIG_ERROR", "failed to load .env app variables", errors.TagInternalServerError)
	}

	if err = envconfig.Process("SWAGGER", &config.Swagger); err != nil {
		return config, errors.New("CONFIG_ERROR", "failed to load .env swagger variables", errors.TagInternalServerError)
	}

	return config, errors.Nil
}
