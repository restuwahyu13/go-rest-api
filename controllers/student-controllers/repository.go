package controllers

// import (
// 	"time"

// 	"github.com/sirupsen/logrus"
// 	"gorm.io/gorm"
// )

// type Repository interface {
// 	FindOneAndCreate(payload EntityAuth) (EntityAuth, error)
// 	Find() (EntityAuth, error)
// 	FindById(id uint) (EntityAuth, error)
// 	FindOneAndDelete(id uint) error
// 	FindOneAndUpdate(id uint, payload EntityAuth) error
// }

// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *repository {
// 	return &repository{db: db}
// }

// func (r *repository) FindOneAndCreate(payload EntityAuth) (EntityAuth, error) {

// 	trx := r.db.Begin()

// 	var student EntityAuth

// 	result := trx.Where("email", payload.Npm).First(&student)
// 	affectedResult := result.RowsAffected
// 	errorResult := result.Error

// 	if affectedResult > 0 {
// 		logrus.Print("npm already taken")
// 		return payload, errorResult
// 	}

// 	if errorResult != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorResult.Error())
// 		return payload, errorResult
// 	}

// 	student.Name = payload.Name
// 	student.Npm = payload.Npm
// 	student.Bid = payload.Bid
// 	student.Fak = payload.Fak
// 	student.CreatedAt = time.Now()

// 	errorCreate := trx.Create(&student).Error

// 	if errorCreate != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCreate.Error())
// 		return payload, errorCreate
// 	}

// 	return student, nil
// }

// func (r *repository) Find() (EntityAuth, error) {
// 	trx := r.db.Begin()

// 	var students EntityAuth

// 	results := trx.Select(&students)
// 	affectedResult := results.RowsAffected
// 	errorResult := results.Error

// 	if affectedResult > 0 {
// 		logrus.Print("student is not exists")
// 		return students, errorResult
// 	}

// 	if errorResult != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorResult.Error())
// 		return students, errorResult
// 	}

// 	return students, nil
// }

// func (r *repository) FindById(id uint) (EntityAuth, error) {
// 	trx := r.db.Begin()

// 	var students EntityAuth

// 	errorResults := trx.Where("ID", id).Select(&students).Error

// 	if errorResults != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorResults.Error())
// 		return students, errorResults
// 	}

// 	return students, nil
// }

// func (r *repository) FindOneAndDelete(id uint) (EntityAuth, error) {
// 	trx := r.db.Begin()

// 	var students EntityAuth

// 	errorCheck := trx.Where("ID", id).Select(&students).Error

// 	if errorCheck != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCheck.Error())
// 		return students, errorCheck
// 	}

// 	errorDelete := trx.Where("ID", id).Delete(&students).Error

// 	if errorDelete != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCheck.Error())
// 		return students, errorCheck
// 	}

// 	return students, nil
// }

// func (r *repository) FindOneAndUpdate(id uint, payload EntityAuth) (EntityAuth, error) {
// 	trx := r.db.Begin()

// 	var student EntityAuth

// 	errorCheck := trx.Where("ID", id).Select(&student).Error

// 	if errorCheck != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCheck.Error())
// 		return student, errorCheck
// 	}

// 	student.Name = payload.Name
// 	student.Npm = payload.Npm
// 	student.Fak = payload.Fak
// 	student.Bid = payload.Bid
// 	student.UpdatedAt = time.Now()

// 	errorUpdate := trx.Where("id", id).Update("id", &student).Error

// 	if errorUpdate != nil {
// 		defer trx.Rollback()
// 		logrus.Fatal(errorCheck.Error())
// 		return student, errorCheck
// 	}

// 	return student, nil
// }
