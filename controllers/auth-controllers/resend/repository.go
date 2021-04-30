package resendAuth

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResend(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "RESEND_NOT_FOUD_404"
		return &users, <-errorCode
	}

	if users.Active {
		errorCode <- "RESEND_ACTIVE_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
