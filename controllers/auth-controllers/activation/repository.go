package activation

import (
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
		return &users, <-errorCode
	}

	db.Select("Active").Where("activation = ?", input.Active).First(&users)

	if users.Active {
		db.Rollback()
		errorCode <- "ACTIVATION_ACTIVE_400"
		return &users, <-errorCode
	}

	data := model.EntityUsers{}
	updateActivationAccount := db.Select("active", "updated_at").Where("email = ?", input.Email).Take(&data).UpdateColumns(map[string]interface{}{
		"email":      input.Email,
		"updated_at": time.Now().Local(),
	})

	if updateActivationAccount.Error != nil {
		db.Rollback()
		errorCode <- "ACTIVATION_ACCOUNT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
