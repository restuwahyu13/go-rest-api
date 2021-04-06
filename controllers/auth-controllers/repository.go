package controllers

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	FindOneAndCreate(payload EntityRegister) (EntityRegister, error)
	Find() (EntityRegister, error)
	FindById(id uint) (EntityRegister, error)
	FindOneAndDelete(id uint) error
	FindOneAndUpdate(id uint, payload EntityRegister) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindOneAndCreate(payload EntityRegister) (EntityRegister, error) {

	trx := r.db.Begin()

	var student EntityRegister

	errorCheck := trx.Where("npm", payload.Npm).First(&student).Error

	if errorCheck != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCheck.Error())
		return payload, errorCheck
	}

	student.Name = payload.Name
	student.Npm = payload.Npm
	student.Bid = payload.Bid
	student.Fak = payload.Fak
	student.CreatedAt = time.Now()

	errorCreate := trx.Create(&student).Error

	if errorCreate != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCreate.Error())
		return payload, errorCreate
	}

	return student, nil
}

func (r *repository) Find() (EntityRegister, error) {
	trx := r.db.Begin()

	var students EntityRegister

	errorResults := trx.Select(&students).Error

	if errorResults != nil {
		defer trx.Rollback()
		logrus.Fatal(errorResults.Error())
		return students, errorResults
	}

	return students, nil
}

func (r *repository) FindById(id uint) (EntityRegister, error) {
	trx := r.db.Begin()

	var students EntityRegister

	errorResults := trx.Where("ID", id).Select(&students).Error

	if errorResults != nil {
		defer trx.Rollback()
		logrus.Fatal(errorResults.Error())
		return students, errorResults
	}

	return students, nil
}

func (r *repository) FindOneAndDelete(id uint) (EntityRegister, error) {
	trx := r.db.Begin()

	var students EntityRegister

	errorCheck := trx.Where("ID", id).Select(&students).Error

	if errorCheck != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCheck.Error())
		return students, errorCheck
	}

	errorDelete := trx.Where("ID", id).Delete(&students).Error

	if errorDelete != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCheck.Error())
		return students, errorCheck
	}

	return students, nil
}

func (r *repository) FindOneAndUpdate(id uint, payload EntityRegister) (EntityRegister, error) {
	trx := r.db.Begin()

	var student EntityRegister

	errorCheck := trx.Where("ID", id).Select(&student).Error

	if errorCheck != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCheck.Error())
		return student, errorCheck
	}

	student.Name = payload.Name
	student.Npm = payload.Npm
	student.Fak = payload.Fak
	student.Bid = payload.Bid
	student.UpdatedAt = time.Now()

	errorUpdate := trx.Where("id", id).Update("id", &student).Error

	if errorUpdate != nil {
		defer trx.Rollback()
		logrus.Fatal(errorCheck.Error())
		return student, errorCheck
	}

	return student, nil
}
