package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers"
	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service auth.Service
}

func NewHandler(service auth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input auth.InputRegister

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	_, errRegister := h.service.RegisterService(&input)

	if errRegister == "REGISTER_CONFLICT_409" {
		response := utils.APIResponse("Email already exist", http.StatusConflict, http.MethodPost, nil)
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if errRegister == "REGISTER_FAILED_403" {
		response := utils.APIResponse("Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := utils.APIResponse("Register new account successfully", http.StatusOK, http.MethodPost, nil)
	ctx.JSON(http.StatusOK, response)
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input auth.InputLogin

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	_, errLogin := h.service.LoginService(&input)

	if errLogin == "LOGIN_NOT_FOUND_404" {
		response := utils.APIResponse("User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if errLogin == "LOGIN_NOT_ACTIVE_403" {
		response := utils.APIResponse("User account is not active", http.StatusForbidden, http.MethodPost, nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	if errLogin == "LOGIN_WRONG_PASSWORD_403" {
		response := utils.APIResponse("Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := utils.APIResponse("Login successfully", http.StatusOK, http.MethodPost, nil)
	ctx.JSON(http.StatusOK, response)
}
