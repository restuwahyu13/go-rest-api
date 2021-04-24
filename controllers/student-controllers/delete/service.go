package createStudent

import (
	"time"

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

	var students model.EntityStudent

	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid
	students.CreatedAt = time.Now().Local()

	resultCreateStudent, errCreateStudent := s.repository.CreateStudentRepository(&students)

	return resultCreateStudent, errCreateStudent
}
