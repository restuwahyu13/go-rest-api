package handlerUpdateStudent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	updateStudent "github.com/restuwahyu13/gin-rest-api/controllers/student-controllers/update"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service updateStudent.Service
}

func NewHandlerUpdateStudent(service updateStudent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateStudentHandler(ctx *gin.Context) {

	var input updateStudent.InputUpdateStudent
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		id := ctx.Param("id")
		toUinteger, _ := strconv.ParseUint(id, 32, 32)

		input.ID = toUinteger
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
}
