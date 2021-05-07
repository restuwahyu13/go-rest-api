package handlerDeleteStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	deleteStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/delete"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
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

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
			gpc.ErrorMetaConfig{
				Tag:     "uuid",
				Field:   "ID",
				Message: "params must be uuid format",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodDelete, errResponse)
		return
	}

	_, errDeleteStudent := h.service.DeleteStudentService(&input)

	switch errDeleteStudent {

	case "DELETE_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Student data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case "DELETE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Delete student data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		util.APIResponse(ctx, "Delete student data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}
