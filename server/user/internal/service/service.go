package service

import "qezde/user/internal/domain/user"

type Configuration func(s *Service) error

type Service struct {
	userRepository user.Service
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithUserRepository(userRepository user.Repository) Configuration {
	return func(s *Service) error {
		s.userRepository = NewUserService(userRepository)
		return nil
	}
}
