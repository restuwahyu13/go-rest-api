package handlerResend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/forgot"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service forgot.Service
}

func NewHandlerForgot(service forgot.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ForgotHandler(ctx *gin.Context) {
	var input forgot.InputForgot

	ctx.ShouldBindJSON(&input)

	forgotResult, errForgot := h.service.ForgotService(&input)

	if errForgot == "FORGOT_NOT_FOUD_404" {
		util.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodGet, nil)
		return
	}

	if errForgot == "FORGOT_NOT_ACTIVE_400" {
		util.APIResponse(ctx, "User account is not active", http.StatusNotFound, http.MethodGet, nil)
		return
	}

	secretKey := util.GodotEnv("JWT_SECRET")
	accessTokenData := map[string]interface{}{"id": forgotResult.ID, "email": forgotResult.Email}
	accessToken, errToken := util.Sign(accessTokenData, secretKey, 5)

	if errToken != nil {
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	_, errorEmail := util.SendGridMail(forgotResult.Fullname, forgotResult.Email, "Reset Password", "template_reset", accessToken)

	if errorEmail != nil {
		util.APIResponse(ctx, "Sending email reset password failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Forgot password successfully", http.StatusNotFound, http.MethodPost, nil)
}
