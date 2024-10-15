package service

import (
	"github.com/google/uuid"
	"qezde/user/db/postgres/sqlc"
	"qezde/user/internal/domain/user"
)

type UserService struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s UserService) GetAllUsers() (dest []user.DTO, err error) {
	data, err := s.repository.GetAllUsers()
	if err != nil {
		return
	}

	dest, err = user.ConvertEntityToDTOs(data)
	return
}

func (s UserService) GetUserById(id uuid.UUID) (dest user.DTO, err error) {
	data, err := s.repository.GetUserById(id)
	if err != nil {
		return
	}

	dest, err = user.ConvertEntityToDTO(data)
	return
}

func (s UserService) CreateUser(params sqlc.CreateUserParams) (dest user.DTO, err error) {
	data, err := s.repository.CreateUser(params)
	if err != nil {
		return
	}

	dest, err = user.ConvertEntityToDTO(data)
	return
}

func (s UserService) UpdateUser(params sqlc.UpdateUserParams) (dest user.DTO, err error) {
	data, err := s.repository.UpdateUser(params)
	if err != nil {
		return
	}

	dest, err = user.ConvertEntityToDTO(data)
	return
}
