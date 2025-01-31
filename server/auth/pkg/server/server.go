package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"qezde/auth/internal/config"
	"syscall"
	"time"
)

type Server struct {
	http *http.Server
	ctx  context.Context
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

func WithHTTPServer(config config.Config, handler http.Handler, ctx context.Context) Configuration {
	return func(s *Server) (err error) {
		s.http = &http.Server{
			Addr:    ":" + config.App.Port,
			Handler: handler,
		}
		s.ctx = ctx
		return nil
	}
}

func (s *Server) Run() (err error) {
	go func() {
		if err = s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(s.ctx, 10*time.Second)
	defer cancel()

	if err = s.http.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %+v", err)
		return
	}

	return
}
