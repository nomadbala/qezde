package service

import (
	"html/template"
	"qezde/notification/internal/config"
	"qezde/notification/pkg/errors"
	"qezde/notification/pkg/resend"
	"sync"
)

type Configuration func(s *Service) errors.Error

type TemplateCache struct {
	sync.RWMutex
	templates map[string]*template.Template
}

type Service struct {
	Resend        *resend.Client
	templateCache TemplateCache
}

func New(configs ...Configuration) (s *Service, err errors.Error) {
	s = &Service{
		templateCache: TemplateCache{
			templates: make(map[string]*template.Template),
		},
	}

	for _, config := range configs {
		if err = config(s); err != errors.Nil {
			return s, errors.New("SERVICE_ERROR", "failed while applying configs to service", errors.TagInternalServerError)
		}
	}

	return s, errors.Nil
}

func WithResendService(config config.ResendConfig) Configuration {
	return func(s *Service) (err errors.Error) {
		if config.APIKey == "" || config.Subject == "" || config.Sender == "" {
			return errors.New("SERVICE_ERROR", "missing required configuration parameters", errors.TagInternalServerError)
		}

		s.Resend = resend.New(config)

		return errors.Nil
	}
}
