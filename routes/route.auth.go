package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers"
)

func initAuthRouter(r *gin.Engine) {
	r.POST("/login", controller.LoginController)
	r.POST("/register", controller.RegisterController)
	return
}
