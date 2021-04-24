package forgotAuth

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
)

type Service interface {
	ForgotService(input *InputForgot) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceForgot(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ForgotService(input *InputForgot) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Email: input.Email,
	}

	resultRegister, errRegister := s.repository.ForgotRepository(&users)

	return resultRegister, errRegister
}
