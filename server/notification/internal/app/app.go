package app

import (
	"log"
	"os"
	"os/signal"
	"qezde/notification/docs"
	"qezde/notification/internal/config"
	"qezde/notification/internal/handler"
	"qezde/notification/internal/service"
	"qezde/notification/pkg/errors"
	"qezde/notification/pkg/server"
	"syscall"
)

func Run() {
	configs, err := config.New()
	if err != errors.Nil {
		log.Fatal(err.Error())
		return
	}

	docs.SwaggerInfo.BasePath = configs.Swagger.BasePath

	services, err := service.New(
		service.WithResendService(configs.Resend),
	)
	if err != errors.Nil {
		log.Fatal(err.Error())
		return
	}

	handlers, err := handler.New(
		handler.WithHTTPHandler(),
		handler.WithResendService(services),
	)
	if err != errors.Nil {
		log.Fatal(err.Error())
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
