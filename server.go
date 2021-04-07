package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/utils"
)

func main() {

	// load dotenv file
	DATABASE_URL := utils.GodotEnv("DATABASE_URL")
	PORT := utils.GodotEnv("PORT")

	// init app
	router := gin.Default()

	// init all route
	// route.InitAuthRoutes(router)
	// route.InitStudentRoutes(router)

	// run server
	router.Run(":3000")
}
