package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qezde/api-gateway/internal/config"
)

type Server struct {
	fiber *fiber.App
	ctx   context.Context
}

type Configuration func(s *Server) error

func New(configs ...Configuration) (r *Server, err error) {
	r = &Server{}
	for _, cfg := range configs {
		if err = cfg(r); err != nil {
			return
		}
	}
	return
}

func WithHTTPServer(handler *fiber.App) Configuration {
	return func(s *Server) (err error) {
		s.fiber = handler
		return
	}
}

func (s *Server) Run(config config.Config) (err error) {
	go func() {
		if err = s.fiber.Listen(":" + config.APP.Port); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Контекст с таймаутом для завершения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Останавливаем сервер Fiber
	if err = s.fiber.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %+v", err)
		return
	}

	return
}
