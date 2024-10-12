package app

import (
	"log"

	"github.com/qezde/api-gateway/internal/config"
	"github.com/qezde/api-gateway/internal/handler"
	"github.com/qezde/api-gateway/pkg/server"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		log.Fatalf("error occurred while loading configs", err)
	}

	handlers, err := handler.New(
		handler.Dependencies{
			Configs: configs,
		},
		handler.WithHTTPHandler())
	if err != nil {
		log.Fatalf("error occurred while initializing handlers", err)
	}

	servers, err := server.New(server.WithHTTPServer(handlers.HTTP, configs))
	if err != nil {
		log.Fatalf("error occurred while initializing server", err)
	}

	if err := servers.Run(); err != nil {
		log.Fatalf("error occurred while running server", err)
	}
}
