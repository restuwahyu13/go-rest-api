package main

import (
	"github.com/gin-gonic/gin"
	route "github.com/restuwahyu13/gin-rest-api/routes"
)

func main() {
	// init app
	router := gin.Default()

	// init all route
	route.InitAuthRoutes(router)
	route.InitStudentRoutes(router)
	route.InitTeacherRoutes(router)

	// run server
	router.Run(":3000")
}
