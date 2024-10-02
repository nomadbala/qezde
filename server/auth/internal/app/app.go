package app

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/service"
	"auth/pkg/server"
	"context"
	"log"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		log.Fatalf("error occurred while loading configs", err)
	}

	ctx := context.Background()

	services, err := service.New(
		service.Dependencies{Config: configs}, service.WithAuthenticationService(),
	)
	if err != nil {
		log.Fatalf("error occurred while initializing services", err)
	}

	handlers, err := handler.New(
		handler.Dependencies{
			Configs:               configs,
			AuthenticationService: services.AuthenticationService,
		},
		handler.WithHTTPHandler(),
	)
	if err != nil {
		log.Fatalf("error occurred while initializing handlers", err)
	}

	servers, err := server.New(server.WithHTTPServer(configs, handlers.HTTP, ctx))
	if err != nil {
		log.Fatalf("error occurred while initializing servers", err)
	}

	if err := servers.Run(); err != nil {
		log.Fatalf("error occurred while initializing servers", err)
	}
}
