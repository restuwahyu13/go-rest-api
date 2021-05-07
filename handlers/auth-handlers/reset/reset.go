package handlerReset

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resetAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/reset"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service resetAuth.Service
}

func NewHandlerReset(service resetAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResetHandler(ctx *gin.Context) {

	var input resetAuth.InputReset
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
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Password",
				Message: "password is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Password",
				Message: "password minimum must be 8 character",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Cpassword",
				Message: "cpassword is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Cpassword",
				Message: "cpassword minimum must be 8 character",
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

	if input.Cpassword != input.Password {
		util.APIResponse(ctx, "Confirm Password is not match with Password", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := util.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errReset := h.service.ResetService(&input)

	switch errReset {

	case "RESET_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not exist", http.StatusNotFound, http.MethodPost, nil)
		return

	case "ACCOUNT_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "RESET_PASSWORD_FAILED_403":
		util.APIResponse(ctx, "Change new password failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Change new password successfully", http.StatusOK, http.MethodPost, nil)
	}
}
