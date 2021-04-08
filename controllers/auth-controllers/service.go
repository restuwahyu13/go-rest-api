package auth

import (
	"time"

	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type Service interface {
	RegisterService(input InputRegister) (EntityUsers, error)
	// LoginService(payload InputLogin) (EntityUsers error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input InputRegister) (EntityUsers, error) {

	users := EntityUsers{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(utils.HashPassword(input.Password)),
		CreatedAt: time.Now(),
	}

	result, err := s.repository.RegisterRepository(users)

	if err != nil {
		logrus.Fatal(err.Error())
		return result, err
	}

	return result, err
}

// func (ctx *service) LoginService(payload InputLogin) (EntityUsers, error) {

// 	user := EntityUsers{
// 		Email:    payload.Email,
// 		Password: payload.Password,
// 	}

// 	result, err := ctx.service.LoginRepository(&user)

// 	if err != nil {
// 		logrus.Fatal(err.Error())
// 		return result, err
// 	}

// 	errCompare := utils.ComparePassword(result.Password, payload.Password)
// 	if errCompare != nil {
// 		logrus.Fatal("password is not match")
// 		return result, errCompare
// 	}
// }
