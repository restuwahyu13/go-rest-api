package handlerLogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	util "github.com/restuwahyu13/gin-rest-api/utils"
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

	resultLogin, errLogin := h.service.LoginService(&input)

	if errLogin == "LOGIN_NOT_FOUND_404" {
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	if errLogin == "LOGIN_NOT_ACTIVE_403" {
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	if errLogin == "LOGIN_WRONG_PASSWORD_403" {
		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	secretKey := util.GodotEnv("JWT_SECRET")
	accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
	accessToken, errToken := util.Sign(accessTokenData, secretKey, 5)

	if errToken != nil {
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
}
