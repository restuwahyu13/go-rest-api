package updateStudent

import (
	"time"

	model "github.com/restuwahyu13/gin-rest-api/models"
)

type Service interface {
	UpdateStudentService(input *InputUpdateStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateStudentService(input *InputUpdateStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent

	students.ID = input.ID
	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid
	students.UpdatedAt = time.Now().Local()

	resultUpdateStudent, errUpdateStudent := s.repository.UpdateStudentRepository(&students)

	return resultUpdateStudent, errUpdateStudent
}
