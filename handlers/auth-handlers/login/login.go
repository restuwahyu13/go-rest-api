package handlerLogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	loginAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.InputLogin

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		resultLogin, errLogin := h.service.LoginService(&input)

		switch errLogin {

		case "LOGIN_NOT_FOUND_404":
			util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)

		case "LOGIN_NOT_ACTIVE_403":
			util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)

		case "LOGIN_WRONG_PASSWORD_403":
			util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)

		default:
			accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
			accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 5)

			if errToken != nil {
				util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
		}
	}
}
