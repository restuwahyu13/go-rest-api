package handlerResultStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resultStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/result"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service resultStudent.Service
}

func NewHandlerResultStudent(service resultStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultStudentHandler(ctx *gin.Context) {

	var input resultStudent.InputResultStudent
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
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	resultStudent, errResultStudent := h.service.ResultStudentService(&input)

	switch errResultStudent {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Student data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Result Student data successfully", http.StatusOK, http.MethodGet, resultStudent)
	}
}
