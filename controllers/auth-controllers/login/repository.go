package loginAuth

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

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
