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
		Server     ServerConfig
		CORS       CORSConfig
		API        APIConfig
		Middleware MiddlewareConfig
	}

	ServerConfig struct {
		Port           string
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}

	CORSConfig struct {
		AllowedOrigins   []string
		AllowedMethods   []string
		AllowedHeaders   []string
		ExposedHeaders   []string
		AllowCredentials bool
	}

	APIConfig struct {
		Auth string
	}

	MiddlewareConfig struct {
		SigningKey string
	}
)

// Server
const (
	DefaultPort    = "8080"
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
)

// CORS
var (
	AllowOrigins  = []string{"*"}
	AllowMethods  = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	AllowHeaders  = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"}
	ExposeHeaders = []string{"Content-Length"}
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

	config.CORS = CORSConfig{
		AllowedOrigins:   AllowOrigins,
		AllowedMethods:   AllowMethods,
		AllowedHeaders:   AllowHeaders,
		ExposedHeaders:   ExposeHeaders,
		AllowCredentials: true,
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
