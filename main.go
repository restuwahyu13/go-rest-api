package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/restuwahyu13/gin-rest-api/configs"
	route "github.com/restuwahyu13/gin-rest-api/routes"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func main() {
	/*
		@description Setup Database Connection
	*/
	db := config.Connection()
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
	router.Run(":" + util.GodotEnv("PORT"))
}
