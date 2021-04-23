package resend

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
	db := r.db.Begin()
	errorCode := make(chan string, 1)

	users := model.EntityUsers{
		Email: input.Email,
	}

	checkUserAccount := db.Select("*").Where("email = ?", input.Email).First(&users).RowsAffected

	if checkUserAccount < 1 {
		db.Rollback()
		errorCode <- "RESEND_NOT_FOUD_404"
		return &users, <-errorCode
	}

	if !users.Active {
		db.Rollback()
		errorCode <- "RESEND_NOT_ACTIVE_400"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
