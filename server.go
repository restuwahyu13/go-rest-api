package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controller"
)

func main() {
	router := gin.Default()
	router.GET("/login", controller.LoginController)
	router.GET("/register", controller.RegisterController)

	router.Run(":3000")
}
