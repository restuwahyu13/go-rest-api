package handlerResend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/resend"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service resend.Service
}

func NewHandlerResend(service resend.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResendHandler(ctx *gin.Context) {
	var input resend.InputResend

	ctx.ShouldBindJSON(&input)

	resendResult, errResend := h.service.ResendService(&input)

	switch errResend {

	case "RESEND_NOT_FOUD_404":
		util.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodGet, nil)
		return

	case "RESEND_NOT_ACTIVE_400":
		util.APIResponse(ctx, "User account is not active", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		secretKey := util.GodotEnv("JWT_SECRET")
		accessTokenData := map[string]interface{}{"id": resendResult.ID, "email": resendResult.Email}
		accessToken, errToken := util.Sign(accessTokenData, secretKey, 5)

		if errToken != nil {
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errorEmail := util.SendGridMail(resendResult.Fullname, resendResult.Email, "Resend New Activation", "template_resend", accessToken)

		if errorEmail != nil {
			util.APIResponse(ctx, "Sending email resend activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Resend new activation token successfully", http.StatusNotFound, http.MethodPost, nil)
	}
}
