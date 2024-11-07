package app

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"qezde/user/internal/config"
	"qezde/user/internal/handler"
	"qezde/user/internal/repository"
	"qezde/user/internal/service"
	logger "qezde/user/pkg/log"
	"qezde/user/pkg/server"
	"syscall"
)

func Run() {
	logger.InitLogger()
	defer func(Log *zap.SugaredLogger) {
		err := Log.Sync()
		if err != nil {
			return
		}
	}(logger.Log)

	configs, err := config.New()
	if err != nil {
		logger.Log.Error("ERR_INIT_CONFIG", zap.Error(err))
		return
	}

	repositories, err := repository.New(repository.WithPostgresStore(configs.Database.DSN))
	if err != nil {
		logger.Log.Error("ERR_INIT_REPOS", zap.Error(err))
		return
	}

	userService, err := service.New(service.WithUserRepository(repositories.User))
	if err != nil {
		logger.Log.Error("ERR_INIT_USER_SERVICE", zap.Error(err))
		return
	}

	handlers, err := handler.New(handler.Dependencies{
		Configs:     configs,
		UserService: userService,
		Ctx:         context.Background(),
	}, handler.WithHTTPHandler())
	if err != nil {
		logger.Log.Error("ERR_INIT_HANDLER", zap.Error(err))
		return
	}

	servers := server.NewServer(configs, handlers.Router)

	servers.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := servers.Shutdown(); err != nil {
		logger.Log.Error("ERR_SERVER_FORCED_SHUTDOWN", zap.Error(err))
		return
	}
}
