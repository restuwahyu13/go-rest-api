package forgot

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
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
	trx := r.db.Begin()
	errorCode := make(chan string, 1)

	users := model.EntityUsers{
		Email: input.Email,
	}

	checkUserAccount := trx.Where("email = ?", input.Email).First(&users).RowsAffected

	if checkUserAccount < 1 {
		errorCode <- "FORGOT_NOT_FOUD_404"
	}

	if !users.Active {
		errorCode <- "FORGOT_NOT_ACTIVE_400"
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
