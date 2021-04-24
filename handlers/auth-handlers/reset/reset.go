package handlerReset

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resetAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/reset"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service resetAuth.Service
}

func NewHandlerReset(service resetAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResetHandler(ctx *gin.Context) {
	var input resetAuth.InputReset
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	}

	token := ctx.Param("token")
	resultToken, errToken := util.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
	}

	if input.Cpassword != input.Password {
		util.APIResponse(ctx, "Confirm Password is not match with Password", http.StatusBadRequest, http.MethodPost, nil)
	}

	result := util.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errReset := h.service.ResetService(&input)

	switch errReset {

	case "RESET_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not exists", http.StatusNotFound, http.MethodPost, nil)

	case "ACCOUNT_NOT_ACTIVE_404":
		util.APIResponse(ctx, "User account is not active", http.StatusBadRequest, http.MethodPost, nil)

	case "RESET_PASSWORD_FAILED_403":
		util.APIResponse(ctx, "Change new password failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Change new password successfully", http.StatusOK, http.MethodPost, nil)
	}
}
