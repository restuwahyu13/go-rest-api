package handlerRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registerAuth "github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	util "github.com/restuwahyu13/gin-rest-api/utils"
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

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		resultRegister, errRegister := h.service.RegisterService(&input)

		switch errRegister {

		case "REGISTER_CONFLICT_409":
			util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)

		case "REGISTER_FAILED_403":
			util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)

		default:
			accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
			accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			_, errSendMail := util.SendGridMail(resultRegister.Fullname, resultRegister.Email, "Activation Account", "template_register", accessToken)

			if errSendMail != nil {
				defer logrus.Error(errSendMail.Error())
				util.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			util.APIResponse(ctx, "Register new account successfully", http.StatusOK, http.MethodPost, nil)
		}
	}
}
