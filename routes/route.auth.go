package route

import (
	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers"
	handler "github.com/restuwahyu13/gin-rest-api/handlers"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description Total Controller
	*/
	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authHandlers := handler.NewHandler(authService)

	/**
	@description Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", authHandlers.RegisterHandler)
	groupRoute.POST("/login", authHandlers.LoginHandler)
}
