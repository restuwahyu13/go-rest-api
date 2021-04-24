package createStudent

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent
	db := r.db.Model(&students)
	errorCode := make(chan string, 1)

	checkStudentExist := db.Select("npm").Where("npm ? =", input.Npm).First(&students).RowsAffected

	if checkStudentExist > 0 {
		errorCode <- "CREATE_STUDENT_CONFLICT_409"
		return &students, <-errorCode
	}

	students.Name = input.Name
	students.Npm = input.Npm
	students.Fak = input.Fak
	students.Bid = input.Bid
	students.CreatedAt = input.CreatedAt

	addNewStudent := db.Create(&students).Error
	db.Commit()

	if addNewStudent != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &students, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &students, <-errorCode
}
