package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/utils"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {
		SecretPublicKey := utils.GodotEnv("JWT_SECRET")
		token, err := utils.Verify(ctx, SecretPublicKey)

		errorResponse := UnathorizatedError{
			Status:  "Unathorizated",
			Code:    http.StatusUnauthorized,
			Method:  ctx.Request.Method,
			Message: "accessToken invalid or expired",
		}

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
		} else {
			// global value result
			ctx.Set("user", token.Claims)
			// return to next method if token is exist
			ctx.Next()
		}
	})
}
