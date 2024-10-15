package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"qezde/api-gateway/internal/config"
)

type Server struct {
	server *http.Server
}

func NewServer(configs config.Config, handlers http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":" + configs.APP.Port,
			Handler: handlers,
		},
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error: ", err)
		}
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.server.Shutdown(ctx)
}
