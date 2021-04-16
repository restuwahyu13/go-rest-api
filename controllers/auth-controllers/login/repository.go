package login

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(payload *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(payload *model.EntityUsers) (*model.EntityUsers, string) {
	trx := r.db.Begin()
	errorCode := make(chan string, 2)

	users := model.EntityUsers{
		Email:    payload.Email,
		Password: payload.Password,
	}

	checkUserAccount := trx.Where("email", payload.Email).First(&users).Error

	if checkUserAccount != nil {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, payload.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
