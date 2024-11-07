package app

import (
	"log"
	"os"
	"os/signal"
	"qezde/notification/docs"
	"qezde/notification/internal/config"
	"qezde/notification/internal/handler"
	"qezde/notification/internal/service"
	"qezde/notification/pkg/server"
	"syscall"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	docs.SwaggerInfo.BasePath = configs.Swagger.BasePath

	services, err := service.New(
		service.WithResendService(configs.Resend),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers, err := handler.New(
		handler.WithHTTPHandler(),
		handler.WithResendService(services),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	servers := server.NewServer(configs, handlers.Router)

	servers.Start()

	log.Printf("server started on port: %s", configs.App.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := servers.Shutdown(); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
}
