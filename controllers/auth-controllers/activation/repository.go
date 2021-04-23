package activation

import (
	"fmt"
	"time"

	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ActivationRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryActivation(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ActivationRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	db := r.db.Begin()
	errorCode := make(chan string, 1)

	users := model.EntityUsers{
		Email: input.Email,
	}

	checkUserAccount := db.Select("*").Where("email = ?", input.Email).First(&users).RowsAffected

	if checkUserAccount < 1 {
		db.Rollback()
		errorCode <- "ACTIVATION_NOT_FOUND_404"
	}

	db.Select("Active").Where("activation = ?", input.Active).First(&users)

	if users.Active {
		db.Rollback()
		errorCode <- "ACTIVATION_ACTIVE_400"
	}

	data := model.EntityUsers{Active: input.Active, UpdatedAt: time.Now().Local()}
	updateActivationAccount := db.Select("Email", "Active", "UpdatedAt").Where("email = ?", input.Email).Updates(&data)
	updateActivationAccount.Commit()

	fmt.Println(updateActivationAccount.Error)

	if updateActivationAccount.Error != nil {
		db.Rollback()
		errorCode <- "ACTIVATION_ACCOUNT_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
