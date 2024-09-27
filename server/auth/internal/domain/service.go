package domain

type Service interface {
	SignUp(request RegistrationRequest) (dest *UserDTO, err error)
	SignIn(request LoginRequest) (token string, err error)
}
