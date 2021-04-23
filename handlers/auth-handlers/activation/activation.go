package handlerActivation

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/activation"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type Handler interface {
	ActivationHandler(ctx *gin.Context)
}

type handler struct {
	service activation.Service
}

func NewHandlerActivation(service activation.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ActivationHandler(ctx *gin.Context) {

	token := ctx.Param("token")
	resultToken, errToken := util.VerifyToken(token, util.GodotEnv("JWT_SECRET"))

	if errToken != nil {
		util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	var data map[string]interface{}
	stringify, _ := json.Marshal(&resultToken)
	json.Unmarshal([]byte(stringify), &data)

	email := data["Claims"].(map[string]interface{})["email"].(string)

	input := activation.InputActivation{
		Email:  email,
		Active: true,
	}

	_, errActivation := h.service.ActivationService(&input)

	if errActivation == "ACTIVATION_NOT_FOUND_404" {
		util.APIResponse(ctx, "User account is not exists", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	if errActivation == "ACTIVATION_ACTIVE_400" {
		util.APIResponse(ctx, "User account hash been active please login", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	if errActivation == "ACTIVATION_ACCOUNT_FAILED_403" {
		util.APIResponse(ctx, "Activation account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Activation account success", http.StatusBadRequest, http.MethodPost, nil)
}
