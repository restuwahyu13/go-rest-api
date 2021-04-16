package handlerLogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service login.Service
}

func NewHandlerLogin(service login.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input login.InputLogin

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	_, errLogin := h.service.LoginService(&input)

	if errLogin == "LOGIN_NOT_FOUND_404" {
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	if errLogin == "LOGIN_NOT_ACTIVE_403" {
		utils.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	if errLogin == "LOGIN_WRONG_PASSWORD_403" {
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, nil)
}
