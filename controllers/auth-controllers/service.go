package auth

import (
	"time"

	"github.com/restuwahyu13/gin-rest-api/utils"
)

type Service interface {
	RegisterService(input InputRegister) (EntityUsers, string)
	LoginService(input InputLogin) (EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input InputRegister) (EntityUsers, string) {

	users := EntityUsers{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(utils.HashPassword(input.Password)),
		CreatedAt: time.Now(),
	}

	resultRegister, errRegister := s.repository.RegisterRepository(users)

	return resultRegister, errRegister
}

func (s *service) LoginService(input InputLogin) (EntityUsers, string) {

	user := EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	loginResult, errLogin := s.repository.LoginRepository(user)

	return loginResult, errLogin
}
