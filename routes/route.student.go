package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers"
)

func InitStudentRoutes(r *gin.Engine) {
	studentRoutes := r.Group("api/v1")
	{
		studentRoutes.POST("/student", controller.CreateStudentController)
	}
}
