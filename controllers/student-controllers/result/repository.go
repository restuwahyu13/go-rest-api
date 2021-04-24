package resultStudent

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ResultStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent
	db := r.db.Model(&students)
	errorCode := make(chan string, 1)

	resultStudents := db.Debug().Select("*").Where("id = ?", input.ID).Find(&students)

	if resultStudents.RowsAffected < 1 {
		errorCode <- "RESULT_STUDENT_NOT_FOUND_404"
		return &students, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &students, <-errorCode
}
