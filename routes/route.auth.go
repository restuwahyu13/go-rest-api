package route

import (
	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	handlerLogin "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/login"
	handlerRegister "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/register"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := login.NewRepositoryLogin(db)
	loginService := login.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := register.NewRepositoryRegister(db)
	registerService := register.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
