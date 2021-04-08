package auth

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(payload EntityUsers) (EntityUsers, error)
	// LoginRepository(payload EntityUsers) (EntityUsers, error)
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

func (r *repository) RegisterRepository(payload EntityUsers) (EntityUsers, error) {

	transaction := r.db.Begin()

	users := EntityUsers{
		Fullname:  payload.Fullname,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: payload.CreatedAt,
	}

	errorResult := transaction.Where("email", payload.Email).First(&users).Error

	if errorResult != nil {
		defer transaction.Rollback()
		logrus.Info(errorResult.Error())
		return payload, errorResult
	}

	errorCreate := transaction.Create(&users).Error
	transaction.Commit()

	if errorCreate != nil {
		defer transaction.Rollback()
		logrus.Info(errorCreate.Error())
		return payload, errorCreate
	}

	return payload, nil
}

// func (r *repository) LoginRepository(payload EntityUsers) (EntityUsers, error) {
// 	trx := r.db.Begin()

// 	var users EntityUsers

// 	results := trx.Select(&users)
// 	affectedResult := results.RowsAffected
// 	errorResult := results.Error

// 	if affectedResult < 1 {
// 		logrus.Info("users is not exists")
// 		return users, errorResult
// 	}

// 	if errorResult != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorResult.Error())
// 		return users, errorResult
// 	}

// 	return users, nil
// }

// func (r *repository) FindById(id uint) (EntityUsers, error) {
// 	trx := r.db.Begin()

// 	var users EntityUsers

// 	errorResults := trx.Where("ID", id).Select(&users).Error

// 	if errorResults != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorResults.Error())
// 		return users, errorResults
// 	}

// 	return users, nil
// }

// func (r *repository) FindOneAndDelete(id uint) (EntityUsers, error) {
// 	trx := r.db.Begin()

// 	var users EntityUsers

// 	errorCheck := trx.Where("ID", id).Select(&users).Error

// 	if errorCheck != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	errorDelete := trx.Where("ID", id).Delete(&users).Error

// 	if errorDelete != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	return users, nil
// }

// func (r *repository) FindOneAndUpdate(id uint, payload EntityUsers) (EntityUsers, error) {
// 	trx := r.db.Begin()

// 	var users EntityUsers

// 	errorCheck := trx.Where("ID", id).Select(&users).Error

// 	if errorCheck != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	users.Name = payload.Name
// 	users.Npm = payload.Npm
// 	users.Fak = payload.Fak
// 	users.Bid = payload.Bid
// 	users.UpdatedAt = time.Now()

// 	errorUpdate := trx.Where("id", id).Update("id", &users).Error

// 	if errorUpdate != nil {
// 		defer trx.Rollback()
// 		logrus.Info(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	return users, nil
// }
