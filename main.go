package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
		@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/*
		@description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	/*
		@description Setup Server
	*/
	router.Run(":" + util.GodotEnv("PORT"))
}
