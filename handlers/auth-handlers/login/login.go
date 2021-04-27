package handlerLogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	loginAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/login"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.InputLogin

	errs := ctx.ShouldBindJSON(&input)

	if errs != nil {
		defer logrus.Error(errs.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	} else {
		errValidator := util.GoValidator(input)
		errResponse := util.ValidatorErrorResponse(errValidator)

		ctx.JSON(http.StatusBadRequest, errResponse)

		// if errValidator["Email"] == "" {
		// 	defer logrus.Error(errValidator)
		// 	util.ValidatorErrorResponse(ctx, "Email is required", http.StatusBadRequest, http.MethodPost, "email", input.Email)
		// 	return
		// }

		// if errValidator["Password"] == "" {
		// 	defer logrus.Error(errValidator)
		// 	util.ValidatorErrorResponse(ctx, "Password is required", http.StatusBadRequest, http.MethodPost, "password", input.Password)
		// 	return
		// }

		// if errValidator["Email"] == input.Email {
		// 	defer logrus.Error(errValidator)
		// 	util.ValidatorErrorResponse(ctx, "Email is not valid", http.StatusBadRequest, http.MethodPost, "email", input.Email)
		// 	return
		// }

		// if errValidator["Password"] == input.Password {
		// 	defer logrus.Error(errValidator)
		// 	util.ValidatorErrorResponse(ctx, "Email is not valid", http.StatusBadRequest, http.MethodPost, "password", input.Password)
		// 	return
		// }

	// 	resultLogin, errLogin := h.service.LoginService(&input)

	// 	switch errLogin {

	// 	case "LOGIN_NOT_FOUND_404":
	// 		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
	// 		return

	// 	case "LOGIN_NOT_ACTIVE_403":
	// 		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
	// 		return

	// 	case "LOGIN_WRONG_PASSWORD_403":
	// 		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
	// 		return

	// 	default:
	// 		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
	// 		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 86400)

	// 		if errToken != nil {
	// 			defer logrus.Error(errToken.Error())
	// 			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
	// 			return
	// 		}

	// 		util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	// 	}
	}
}
