package service

import (
	"auth/internal/config"
)

type Dependencies struct {
	Config config.Config
}

type Service struct {
	dependencies          Dependencies
	AuthenticationService *AuthenticationService
}

type Configuration func(s *Service) error

func New(d Dependencies, configs ...Configuration) (s *Service, err error) {
	s = &Service{
		dependencies: d,
	}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithAuthenticationService() Configuration {
	return func(s *Service) error {
		s.AuthenticationService = NewAuthenticationService(s.dependencies.Config)
		return nil
	}
}
