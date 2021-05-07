package handlerActivation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	activationAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/activation"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service activationAuth.Service
}

func NewHandlerActivation(service activationAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ActivationHandler(ctx *gin.Context) {

	var input activationAuth.InputActivation

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Token",
				Message: "accessToken is required on params",
			},
		},
	}

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	token := ctx.Param("token")
	resultToken, errToken := util.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := util.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errActivation := h.service.ActivationService(&input)

	switch errActivation {

	case "ACTIVATION_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not exist", http.StatusNotFound, http.MethodPost, nil)
		return

	case "ACTIVATION_ACTIVE_400":
		util.APIResponse(ctx, "User account hash been active please login", http.StatusBadRequest, http.MethodPost, nil)
		return

	case "ACTIVATION_ACCOUNT_FAILED_403":
		util.APIResponse(ctx, "Activation account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Activation account success", http.StatusOK, http.MethodPost, nil)
	}
}
