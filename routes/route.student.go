package route

import (
	"github.com/gin-gonic/gin"
	createStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/create"
	resultStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/result"
	resultsStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/results"
	handlerCreateStudent "github.com/restuwahyu13/gin-rest-api/handlers/student-handlers/create"
	handlerResultStudent "github.com/restuwahyu13/gin-rest-api/handlers/student-handlers/result"
	handlerResultsStudent "github.com/restuwahyu13/gin-rest-api/handlers/student-handlers/results"
	middleware "github.com/restuwahyu13/gin-rest-api/middlewares"
	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.PUT("/student", createStudentHandler.CreateStudentHandler)
}
