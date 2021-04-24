package activationAuth

import model "github.com/restuwahyu13/gin-rest-api/models"

type Service interface {
	ActivationService(input *InputActivation) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceActivation(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ActivationService(input *InputActivation) (*model.EntityUsers, string) {
	users := model.EntityUsers{
		Email:  input.Email,
		Active: input.Active,
	}

	activationResult, activationError := s.repository.ActivationRepository(&users)

	return activationResult, activationError
}
