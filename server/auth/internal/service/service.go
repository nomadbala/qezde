package service

import (
	"auth/internal/config"
	"auth/internal/domain"
)

type Dependencies struct {
	config config.Config
}

type Service struct {
	dependencies          Dependencies
	authenticationService domain.Service
}

type Configuration func(s *Service) error

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithConfig(c config.Config) Configuration {
	return func(s *Service) error {
		s.dependencies.config = c
		s.authenticationService = NewAuthenticationService(c)
		return nil
	}
}
