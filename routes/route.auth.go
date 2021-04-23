package route

import (
	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/activation"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/resend"
	handlerActivation "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/activation"
	handlerLogin "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/login"
	handlerRegister "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/register"
	handlerResend "github.com/restuwahyu13/gin-rest-api/handlers/auth-handlers/resend"
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

	resendRepository := resend.NewRepositoryResend(db)
	resendService := resend.NewServiceResend(resendRepository)
	resendHandler := handlerResend.NewHandlerResend(resendService)

	activationRepository := activation.NewRepositoryActivation(db)
	activationService := activation.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend", resendHandler.ResendHandler)

}
