package service

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"qezde/auth/internal/config"
	"qezde/auth/internal/domain"
	auth2 "qezde/protogen/auth"
)

type AuthenticationService struct {
	config config.Config
	client auth2.AuthServiceClient
}

func NewAuthenticationService(config config.Config) *AuthenticationService {
	return &AuthenticationService{config: config}
}

func (s AuthenticationService) SignUp(request domain.RegistrationRequest) (dest *domain.UserDTO, err error) {
	//salt, err := generateSalt()
	//if err != nil {
	//	return
	//}

	//hashedPassword, err := hashPassword(request.Password, salt)
	//if err != nil {
	//	return
	//}

	//createUserRequest := domain.CreateUserRequest{
	//	Username:     request.Username,
	//	PasswordHash: hashedPassword,
	//	Salt:         salt,
	//	Email:        request.Email,
	//}

	//TODO: Доделать логику регистрации. Сделать HTTP.POST("users/...")
	panic("implement me")
}

func (s AuthenticationService) SignIn(request domain.LoginRequest) (token string, err error) {
	//TODO implement me
	panic("implement me")
}

func generateSalt() (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}

func hashPassword(password, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
