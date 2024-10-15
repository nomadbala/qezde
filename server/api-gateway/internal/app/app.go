package app

import (
	"log"
	"os"
	"os/signal"
	"qezde/api-gateway/pkg/server"
	"syscall"

	"qezde/api-gateway/internal/config"
	"qezde/api-gateway/internal/handler"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		log.Fatal("error occurred while loading configs", err)
	}

	handlers := handler.New(
		handler.Dependencies{
			Configs: configs,
		},
	)

	servers := server.NewServer(configs, handlers)

	servers.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := servers.Shutdown(); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
}
