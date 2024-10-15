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
		log.Fatal("error occurred while loading configs", err)
	}

	handlers, err := handler.New(
		handler.Dependencies{
			Configs: configs,
		},
		handler.WithHTTPHandler())
	if err != nil {
		log.Fatal("error occurred while initializing handlers", err)
	}

	servers, err := server.New(server.WithHTTPServer(handlers.HTTP))
	if err != nil {
		log.Fatal("error occurred while initializing server", err)
	}

	if err := servers.Run(configs); err != nil {
		log.Fatal("error occurred while running server", err)
	}
}
