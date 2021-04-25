package updateStudent

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent
	db := r.db.Model(&students)
	errorCode := make(chan string, 1)

	students.ID = input.ID

	checkStudentId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- "UPDATE_STUDENT_NOT_FOUND_404"
		return &students, <-errorCode
	}

	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid

	updateStudent := db.Debug().Select("name", "npm", "fak", "bid", "updated_at").Where("id = ?", input.ID).Updates(students)

	if updateStudent.Error != nil {
		errorCode <- "UPDATE_STUDENT_FAILED_403"
		return &students, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &students, <-errorCode
}
