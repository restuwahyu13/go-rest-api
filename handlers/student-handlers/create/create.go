package handlerCreateStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/create"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service createStudent.Service
}

func NewHandlerCreateStudent(service createStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateStudentHandler(ctx *gin.Context) {

	var input createStudent.InputCreateStudent
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
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
				Tag:     "numeric",
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
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errCreateStudent := h.service.CreateStudentService(&input)

	switch errCreateStudent {

	case "CREATE_STUDENT_CONFLICT_409":
		util.APIResponse(ctx, "Npm student already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Create new student account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new student account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
