package resetAuth

import model "github.com/restuwahyu13/gin-rest-api/models"

type Service interface {
	ResetService(input *InputReset) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceReset(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResetService(input *InputReset) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
		Active:   input.Active,
	}

	resetResult, errResult := s.repository.ResetRepository(&users)

	return resetResult, errResult
}
