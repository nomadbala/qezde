package server

import (
	"context"
	"errors"
	"gateway/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	server *http.Server
	ctx    context.Context
}

func New(config config.Config, handler http.Handler, ctx context.Context) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr:           ":" + config.Server.Port,
			Handler:        handler,
			MaxHeaderBytes: config.Server.MaxHeaderBytes,
			ReadTimeout:    config.Server.ReadTimeout,
			WriteTimeout:   config.Server.WriteTimeout,
		},
		ctx: ctx,
	}
}

func (s *HttpServer) Run() (err error) {
	go func() {
		if err = s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(s.ctx, 10*time.Second)
	defer cancel()

	if err = s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %+v", err)
		return
	}

	return
}
