package handlerForgot

import (
	"net/http"

	"github.com/gin-gonic/gin"
	forgotAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/forgot"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service forgotAuth.Service
}

func NewHandlerForgot(service forgotAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ForgotHandler(ctx *gin.Context) {

	var input forgotAuth.InputForgot
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

	forgotResult, errForgot := h.service.ForgotService(&input)

	switch errForgot {

	case "FORGOT_NOT_FOUD_404":
		util.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "FORGOT_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "FORGOT_PASSWORD_FAILED_403":
		util.APIResponse(ctx, "Forgot password failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": forgotResult.ID, "email": forgotResult.Email}
		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 5)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errorEmail := util.SendGridMail(forgotResult.Fullname, forgotResult.Email, "Reset Password", "template_reset", accessToken)

		if errorEmail != nil {
			util.APIResponse(ctx, "Sending email reset password failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Forgot password successfully", http.StatusOK, http.MethodPost, nil)
	}
}
