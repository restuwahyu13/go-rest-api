package handlerRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.InputRegister

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	_, errRegister := h.service.RegisterService(&input)

	if errRegister == "REGISTER_CONFLICT_409" {
		utils.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return
	}

	if errRegister == "REGISTER_FAILED_403" {
		utils.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	utils.APIResponse(ctx, "Register new account successfully", http.StatusOK, http.MethodPost, nil)
}
