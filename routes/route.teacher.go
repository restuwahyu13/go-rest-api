package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/restuwahyu13/gin-rest-api/controllers/teacher-controllers"
)

func InitTeacherRoutes(r *gin.Engine) {
	teacherRoutes := r.Group("api/v1")
	{
		teacherRoutes.POST("/teacher", controller.CreateTeacherController)
	}
}
