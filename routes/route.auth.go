package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers"
)

func InitAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("api/v1")
	{
		authRoutes.POST("/login", controller.LoginController)
		authRoutes.POST("/register", controller.RegisterController)
	}
}
