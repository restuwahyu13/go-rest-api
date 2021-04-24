package handlerResend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resendAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/resend"
	util "github.com/restuwahyu13/gin-rest-api/utils"
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

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		resendResult, errResend := h.service.ResendService(&input)

		switch errResend {

		case "RESEND_NOT_FOUD_404":
			util.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodGet, nil)

		case "RESEND_NOT_ACTIVE_400":
			util.APIResponse(ctx, "User account is not active", http.StatusNotFound, http.MethodGet, nil)

		default:
			accessTokenData := map[string]interface{}{"id": resendResult.ID, "email": resendResult.Email}
			accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 60)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			_, errorSendEmail := util.SendGridMail(resendResult.Fullname, resendResult.Email, "Resend New Activation", "template_resend", accessToken)

			if errorSendEmail != nil {
				defer logrus.Error(errorSendEmail.Error())
				util.APIResponse(ctx, "Sending email resend activation failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			util.APIResponse(ctx, "Resend new activation token successfully", http.StatusNotFound, http.MethodPost, nil)
		}
	}
}
