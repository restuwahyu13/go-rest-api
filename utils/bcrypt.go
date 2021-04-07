package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return result
}

func ComparePassword(hashPassword []byte, password string) {
	pw := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashPassword, pw)
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
}
