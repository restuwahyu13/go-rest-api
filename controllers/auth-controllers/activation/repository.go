package activationAuth

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

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "ACTIVATION_NOT_FOUND_404"
		return &users, <-errorCode
	}

	db.Debug().Select("Active").Where("activation = ?", input.Active).Take(&users)

	if users.Active {
		errorCode <- "ACTIVATION_ACTIVE_400"
		return &users, <-errorCode
	}

	users.Active = input.Active
	users.UpdatedAt = time.Now().Local()

	updateActivation := db.Debug().Select("active", "updated_at").Where("email = ?", input.Email).Updates(users)

	if updateActivation.Error != nil {
		errorCode <- "ACTIVATION_ACCOUNT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
