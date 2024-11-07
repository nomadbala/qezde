package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"qezde/user/internal/config"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(configs config.Config, handlers http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":" + configs.App.Port,
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
