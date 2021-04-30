package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	config "github.com/restuwahyu13/gin-rest-api/configs"
	route "github.com/restuwahyu13/gin-rest-api/routes"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func main() {
	/**
	@description Setup Database Connection
	*/
	db := config.Connection()
	/**
	@description Setup Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if util.GodotEnv("GO_ENV") != "production" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	route.InitStudentRoutes(db, router)
	/**
	@description Setup Server
	*/
	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}
