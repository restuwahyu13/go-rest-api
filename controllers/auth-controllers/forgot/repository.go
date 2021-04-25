package forgotAuth

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ForgotRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryForgot(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ForgotRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = util.HashPassword(util.RandStringBytes(20))

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "FORGOT_NOT_FOUD_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "FORGOT_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	changePassword := db.Debug().Select("password", "updated_at").Where("email = ?", input.Email).Updates(users)

	if changePassword.Error != nil {
		errorCode <- "FORGOT_PASSWORD_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
