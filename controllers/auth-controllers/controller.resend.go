package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResendController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl resend",
	})
}
