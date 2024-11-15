package service

import (
	"qezde/notification/internal/config"
	"qezde/notification/pkg/errors"
	"qezde/notification/pkg/resend"
)

type Configuration func(s *Service) errors.Error

type Service struct {
	Resend *resend.Client
}

func New(configs ...Configuration) (s *Service, err errors.Error) {
	s = &Service{}

	for _, config := range configs {
		if err = config(s); err != errors.Nil {
			return s, errors.New("SERVICE_ERROR", "failed while applying configs to service")
		}
	}

	return s, errors.Nil
}

func WithResendService(config config.ResendConfig) Configuration {
	return func(s *Service) (err errors.Error) {
		if config.APIKey == "" || config.Subject == "" || config.Sender == "" {
			return errors.New("SERVICE_ERROR", "missing required configuration parameters")
		}

		s.Resend = resend.New(config)

		return errors.Nil
	}
}
