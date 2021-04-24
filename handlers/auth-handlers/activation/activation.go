package handlerActivation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	activationAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/activation"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service activationAuth.Service
}

func NewHandlerActivation(service activationAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ActivationHandler(ctx *gin.Context) {

	var input activationAuth.InputActivation
	token := ctx.Param("token")
	resultToken, errToken := util.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
	}

	result := util.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errActivation := h.service.ActivationService(&input)

	switch errActivation {

	case "ACTIVATION_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not exists", http.StatusNotFound, http.MethodPost, nil)

	case "ACTIVATION_ACTIVE_400":
		util.APIResponse(ctx, "User account hash been active please login", http.StatusBadRequest, http.MethodPost, nil)

	case "ACTIVATION_ACCOUNT_FAILED_403":
		util.APIResponse(ctx, "Activation account failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Activation account success", http.StatusOK, http.MethodPost, nil)
	}
}
