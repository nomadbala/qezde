package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	App      AppConfig
	Cors     CorsConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Port string
}

type CorsConfig struct {
	AllowedOrigins []string
}

type DatabaseConfig struct {
	DSN string
}

func New() (config Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))

	if err = envconfig.Process("APP", &config.App); err != nil {
		return
	}

	if err = envconfig.Process("CORS", &config.Cors); err != nil {
		return
	}

	if err = envconfig.Process("POSTGRES", &config.Database); err != nil {
		return
	}

	return
}
