package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
	"time"
)

type (
	Config struct {
		Server ServerConfig
		App    AppConfig
	}

	ServerConfig struct {
		Port           string
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}

	AppConfig struct {
		Path    string
		Port    string
		Timeout time.Duration
	}
)

// Server
const (
	DefaultPort    = "8080"
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
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

	config.Server = ServerConfig{
		Port:           DefaultPort,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
	}

	if err = envconfig.Process("APP", &config); err != nil {
		return
	}

	return
}
