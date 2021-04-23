package login

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	db := r.db.Begin()
	errorCode := make(chan string, 1)

	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	checkUserAccount := db.Select("*").Where("email = ?", input.Email).First(&users).Error

	if checkUserAccount != nil {
		db.Rollback()
		errorCode <- "LOGIN_NOT_FOUND_404"
	}

	if !users.Active {
		db.Rollback()
		errorCode <- "LOGIN_NOT_ACTIVE_403"
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		db.Rollback()
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
