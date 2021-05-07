package handlerRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registerAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service registerAuth.Service
}

func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var input registerAuth.InputRegister
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fullname",
				Message: "fullname is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Fullname",
				Message: "fullname must be using lowercase",
			},
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
		},
	}

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.RegisterService(&input)

	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errSendMail := util.SendGridMail(resultRegister.Fullname, resultRegister.Email, "Activation Account", "template_register", accessToken)

		if errSendMail != nil {
			defer logrus.Error(errSendMail.Error())
			util.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
