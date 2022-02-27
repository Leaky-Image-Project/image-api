package service

import (
	"leaky-image-project/chat-api/entity"
)

type AuthService interface {
	Login(username string, password string) bool
}

type authService struct {
	credentialList []entity.User
}

func NewAuthService() AuthService {
	return &authService{
		credentialList: []entity.User{
			{
				Name:     "attacker",
				Password: "attackPass",
			},
			{
				Name:     "victim0",
				Password: "victim0Pass",
			},
			{
				Name:     "victim1",
				Password: "victim1Pass",
			},
		},
	}
}

func (service *authService) Login(username string, password string) bool {
	for _, v := range service.credentialList {
		if v.Name == username && v.Password == password {
			return true
		}
	}

	return false
}
