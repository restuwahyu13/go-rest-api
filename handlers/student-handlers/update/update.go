package handlerUpdateStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	updateStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/update"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service updateStudent.Service
}

func NewHandlerUpdateStudent(service updateStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateStudentHandler(ctx *gin.Context) {

	var input updateStudent.InputUpdateStudent
	input.ID = ctx.Param("id")
	ctx.ShouldBindJSON(&input)

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
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Name",
				Message: "name must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Npm",
				Message: "npm is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "number",
				Field:   "Npm",
				Message: "npm must be number format",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fak",
				Message: "fak is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Fak",
				Message: "fak must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Bid",
				Message: "bid is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Bid",
				Message: "bid must be using lowercase",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	_, errUpdateStudent := h.service.UpdateStudentService(&input)

	switch errUpdateStudent {

	case "UPDATE_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Student data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Update student data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Update student data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}
