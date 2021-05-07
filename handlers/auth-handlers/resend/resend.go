package handlerResend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resendAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/resend"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service resendAuth.Service
}

func NewHandlerResend(service resendAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResendHandler(ctx *gin.Context) {

	var input resendAuth.InputResend
	ctx.ShouldBindJSON(&input)

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
		},
	}

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resendResult, errResend := h.service.ResendService(&input)

	switch errResend {

	case "RESEND_NOT_FOUD_404":
		util.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "RESEND_ACTIVE_403":
		util.APIResponse(ctx, "User account hash been active", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resendResult.ID, "email": resendResult.Email}
		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 5)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errorSendEmail := util.SendGridMail(resendResult.Fullname, resendResult.Email, "Resend New Activation", "template_resend", accessToken)

		if errorSendEmail != nil {
			defer logrus.Error(errorSendEmail.Error())
			util.APIResponse(ctx, "Sending email resend activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Resend new activation token successfully", http.StatusOK, http.MethodPost, nil)
	}
}
