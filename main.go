package main

import (
	"github.com/gin-gonic/gin"
	model "github.com/restuwahyu13/gin-rest-api/models"
	route "github.com/restuwahyu13/gin-rest-api/routes"
	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	/*
		@description Setup Database Connection
	*/
	databaseURI := utils.GodotEnv("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
		return
	}

	if utils.GodotEnv("NODE_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	/*
		@description Setup Router
	*/
	router := gin.Default()
	/*
		@description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	/*
		@description Setup Server
	*/
	router.Run(":" + utils.GodotEnv("PORT"))
}
