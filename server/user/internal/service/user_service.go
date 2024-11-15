package service

import (
	"context"
	"github.com/google/uuid"
	"qezde/user/internal/domain/user"
)

func (s *Service) GetAllUsers(ctx context.Context) (dest []user.DTO, err error) {
	entities, err := s.userRepository.GetAllUsers(ctx)
	if err != nil {
		return
	}

	dest = make([]user.DTO, len(entities))

	for i, u := range entities {
		dest[i] = u.ConvertToDTO()
	}

	return dest, nil
}

func (s *Service) GetUserById(ctx context.Context, id string) (dest user.DTO, err error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return
	}

	entity, err := s.userRepository.GetUserById(ctx, uuid)

	if entity == nil {
		return dest, user.ErrNotFound
	}

	if err != nil {
		return
	}

	dest = entity.ConvertToDTO()

	return
}

func (s *Service) CreateUser(ctx context.Context, request user.CreateUserRequest) (dest user.DTO, err error) {
	entity, err := s.userRepository.CreateUser(ctx, request)
	if err != nil {
		return
	}

	dest = entity.ConvertToDTO()

	return
}

func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, request user.UpdateUserRequest) (dest user.DTO, err error) {
	entity, err := s.userRepository.UpdateUser(ctx, id, request)

	if entity == nil {
		return dest, user.ErrNotFound
	}

	if err != nil {
		return
	}

	dest = entity.ConvertToDTO()

	return
}
