package app

import (
	"context"
	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/pkg/server"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	dependencies := handler.Dependencies{
		Configs: configs,
	}

	handlers, err := handler.New(dependencies, handler.WithHTTPHandler())
	if err != nil {
		panic(err)
	}

	servers, err := server.New(server.WithHTTPServer(configs, handlers.HTTP, ctx))
	if err != nil {
		panic(err)
	}

	if err := servers.Run(); err != nil {
		panic(err)
	}
}
