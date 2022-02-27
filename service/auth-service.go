package service

type AuthService interface {
	Login(username string, password string) bool
}

type authService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewAuthService() AuthService {
	return &authService{
		authorizedUsername: "attacker",
		authorizedPassword: "passAttack",
	}
}

func (service *authService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
