package server

import (
	"context"
	"net/http"
)

type HttpServer struct {
	server *http.Server
	ctx    context.Context
}

type Configuration struct {
}

func New(configuration *Configuration) *HttpServer {
	return &HttpServer{
		server: &http.Server{

		}
	}
}
