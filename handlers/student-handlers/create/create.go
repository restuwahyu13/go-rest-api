package handlerCreateStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/create"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service createStudent.Service
}

func NewHandlerCreateStudent(service createStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateStudentHandler(ctx *gin.Context) {

	var input createStudent.InputCreateStudent
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	}

	_, errCreateStudent := h.service.CreateStudentService(&input)

	switch errCreateStudent {

	case "CREATE_STUDENT_CONFLICT_409":
		util.APIResponse(ctx, "Npm student already exist", http.StatusConflict, http.MethodPost, nil)

	case "CREATE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Create new student account failed", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Create new student account successfully", http.StatusConflict, http.MethodPost, nil)
	}
}
