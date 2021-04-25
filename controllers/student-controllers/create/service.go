package createStudent

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
)

type Service interface {
	CreateStudentService(input *InputCreateStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateStudentService(input *InputCreateStudent) (*model.EntityStudent, string) {

	students := model.EntityStudent{
		Name: input.Name,
		Npm:  input.Npm,
		Fak:  input.Fak,
		Bid:  input.Bid,
	}

	resultCreateStudent, errCreateStudent := s.repository.CreateStudentRepository(&students)

	return resultCreateStudent, errCreateStudent
}
