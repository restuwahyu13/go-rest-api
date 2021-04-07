package controllers

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(payload EntityUsers) (EntityUsers, error)
	LoginRepository(payload EntityUsers) (EntityUsers, error)
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

	trx := r.db.Begin()

	var users EntityUsers

	result := trx.Where("email", payload.Email).First(&users)
	affectedResult := result.RowsAffected
	errorResult := result.Error

	if affectedResult > 0 {
		logrus.Fatal("email already taken")
		return payload, errorResult
	}

	if errorResult != nil {
		defer trx.Rollback()
		logrus.Fatal(errorResult.Error())
		return payload, errorResult
	}

	users.Fullname = payload.Fullname
	users.Email = payload.Email
	users.Password = payload.Password
	users.CreatedAt = time.Now()

	errorCreate := trx.Create(&users).Error

	if errorCreate != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCreate.Error())
		return payload, errorCreate
	}

	return users, nil
}

func (r *repository) LoginRepository(payload EntityUsers) (EntityUsers, error) {
	trx := r.db.Begin()

	var users EntityUsers

	results := trx.Select(&users)
	affectedResult := results.RowsAffected
	errorResult := results.Error

	if affectedResult > 0 {
		logrus.Fatal("users is not exists")
		return users, errorResult
	}

	if errorResult != nil {
		defer trx.Rollback()
		logrus.Fatal(errorResult.Error())
		return users, errorResult
	}

	return users, nil
}

// func (r *repository) FindById(id uint) (EntityUsers, error) {
// 	trx := r.db.Begin()

// 	var users EntityUsers

// 	errorResults := trx.Where("ID", id).Select(&users).Error

// 	if errorResults != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorResults.Error())
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
// 		logrus.Fatal(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	errorDelete := trx.Where("ID", id).Delete(&users).Error

// 	if errorDelete != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCheck.Error())
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
// 		logrus.Fatal(errorCheck.Error())
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
// 		logrus.Fatal(errorCheck.Error())
// 		return users, errorCheck
// 	}

// 	return users, nil
// }
