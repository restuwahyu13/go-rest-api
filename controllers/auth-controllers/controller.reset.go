package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl reset",
	})
}
