package config

import (
	model "github.com/restuwahyu13/gin-rest-api/models"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := util.GodotEnv("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if util.GodotEnv("NODE_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
