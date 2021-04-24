package handlerCreateStudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/create"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
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
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		_, errCreateStudent := h.service.CreateStudentService(&input)

		switch errCreateStudent {

		case "CREATE_STUDENT_CONFLICT_409":
			util.APIResponse(ctx, "Npm student already exist", http.StatusConflict, http.MethodPost, nil)

		case "CREATE_STUDENT_FAILED_403":
			util.APIResponse(ctx, "Create new student account failed", http.StatusForbidden, http.MethodPost, nil)

		default:
			util.APIResponse(ctx, "Create new student account successfully", http.StatusOK, http.MethodPost, nil)
		}
	}
}
