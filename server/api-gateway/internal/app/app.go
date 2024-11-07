package app

import (
	"os"
	"os/signal"
	"qezde/api-gateway/docs"
	logger "qezde/api-gateway/pkg/log"
	"qezde/api-gateway/pkg/server"
	"syscall"

	"go.uber.org/zap"

	"qezde/api-gateway/internal/config"
	"qezde/api-gateway/internal/handler"
)

func Run() {
	logger.InitLogger()
	defer func(Log *zap.SugaredLogger) {
		err := Log.Sync()
		if err != nil {
			return
		}
	}(logger.Log)

	docs.SwaggerInfo.BasePath = "/"

	configs, err := config.New()
	if err != nil {
		logger.Log.Error("error occurred while loading configs", err)
		return
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
		logger.Log.Error("server forced to shutdown: ", err)
		return
	}
}
