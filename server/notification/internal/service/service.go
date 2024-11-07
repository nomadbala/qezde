package service

import (
	"qezde/notification/internal/config"
	"qezde/notification/pkg/resend"
)

type Configuration func(s *Service) error

type Service struct {
	Resend *resend.Client
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, config := range configs {
		if err = config(s); err != nil {
			return
		}
	}

	return
}

func WithResendService(config config.ResendConfig) Configuration {
	return func(s *Service) (err error) {
		s.Resend = resend.New(config)

		return
	}
}
