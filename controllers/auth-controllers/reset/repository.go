package resetAuth

import (
	"time"

	model "github.com/restuwahyu13/gin-rest-api/models"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResetRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryReset(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResetRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password
	users.Active = input.Active

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "RESET_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "ACCOUNT_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	users.Password = util.HashPassword(input.Password)
	users.UpdatedAt = time.Now().Local()

	updateNewPassword := db.Debug().Select("password", "update_at").Where("email = ?", input.Email).Updates(users)

	if updateNewPassword.Error != nil {
		errorCode <- "RESET_PASSWORD_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
