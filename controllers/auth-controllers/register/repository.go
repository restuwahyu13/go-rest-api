package register

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(payload *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(payload *model.EntityUsers) (*model.EntityUsers, string) {

	transaction := r.db.Begin()
	errorCode := make(chan string, 2)

	users := model.EntityUsers{
		Fullname:  payload.Fullname,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: payload.CreatedAt,
	}

	checkUserAccount := transaction.Where("email", payload.Email).First(&users).RowsAffected

	if checkUserAccount > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
	}

	addNewUser := transaction.Create(&users).Error
	transaction.Commit()

	if addNewUser != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
