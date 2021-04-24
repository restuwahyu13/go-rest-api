package handlerReset

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/reset"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service reset.Service
}

func NewHandlerReset(service reset.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResetHandler(ctx *gin.Context) {
	var input reset.InputReset
	ctx.ShouldBindJSON(&input)

	token := ctx.Param("token")
	resultToken, errToken := util.VerifyToken(token, util.GodotEnv("JWT_SECRET"))

	if errToken != nil {
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

	if errReset == "RESET_NOT_FOUND_404" {
		util.APIResponse(ctx, "User account is not exists", http.StatusNotFound, http.MethodPost, nil)
		return
	}

	if errReset == "ACCOUNT_NOT_ACTIVE_404" {
		util.APIResponse(ctx, "User account is not active", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	if errReset == "RESET_PASSWORD_FAILED_403" {
		util.APIResponse(ctx, "Change new password failed", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Change new password successfully", http.StatusOK, http.MethodPost, nil)
}
