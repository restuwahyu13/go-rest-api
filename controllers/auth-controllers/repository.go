package auth

import (
	"github.com/restuwahyu13/gin-rest-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(payload EntityUsers) (EntityUsers, string)
	LoginRepository(payload EntityUsers) (EntityUsers, string)
	// ActivationRepository(token string, payload EntityUsers) (EntityUsers, error)
	// ForgotRepository(payload EntityUsers) (EntityUsers, error)
	// ResendRepository(payload EntityUsers) (EntityUsers, error)
	// ResetRepository(token, password, cpassword string) (EntityUsers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(payload EntityUsers) (EntityUsers, string) {

	transaction := r.db.Begin()
	errorCode := make(chan string, 2)

	users := EntityUsers{
		Fullname:  payload.Fullname,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: payload.CreatedAt,
	}

	checkUserAccount := transaction.Where("email", payload.Email).First(&users).RowsAffected

	if checkUserAccount > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
	}

	addNewUser := transaction.Create(&users).Error
	transaction.Commit()

	if addNewUser != nil {
		errorCode <- "REGISTER_FAILED_403"
		return users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return users, <-errorCode
}

func (r *repository) LoginRepository(payload EntityUsers) (EntityUsers, string) {
	trx := r.db.Begin()
	errorCode := make(chan string, 2)

	users := EntityUsers{
		Email:    payload.Email,
		Password: payload.Password,
	}

	checkUserAccount := trx.Where("email", payload.Email).First(&users).Error

	if checkUserAccount != nil {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return users, <-errorCode
	}

	comparePassword := utils.ComparePassword(users.Password, payload.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return users, <-errorCode
}
