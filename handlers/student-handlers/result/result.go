package handlerResultStudent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	resultStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/result"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service resultStudent.Service
}

func NewHandlerResultStudent(service resultStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultStudentHandler(ctx *gin.Context) {

	var input resultStudent.InputResultStudent
	id := ctx.Param("id")
	toUinteger, _ := strconv.ParseUint(id, 32, 32)

	input.ID = toUinteger

	resultStudent, errResultStudent := h.service.ResultStudentService(&input)

	switch errResultStudent {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Student data is not exist or deleted", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Result Student data successfully", http.StatusOK, http.MethodPost, resultStudent)
	}
}
