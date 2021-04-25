package handlerDeleteStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	deleteStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/delete"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service deleteStudent.Service
}

func NewHandlerDeleteStudent(service deleteStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteStudentHandler(ctx *gin.Context) {

	var input deleteStudent.InputDeleteStudent
	input.ID = ctx.Param("id")

	_, errDeleteStudent := h.service.DeleteStudentService(&input)

	switch errDeleteStudent {

	case "DELETE_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Student data is not exist or deleted", http.StatusForbidden, http.MethodPost, nil)

	case "DELETE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Delete student data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Delete student data successfully", http.StatusOK, http.MethodPost, nil)
	}
}
