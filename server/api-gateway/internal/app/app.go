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

	handlers := handler.New(configs)

	servers := server.New(configs, handlers, ctx)

	if err := servers.Run(); err != nil {
		panic(err)
	}
}
