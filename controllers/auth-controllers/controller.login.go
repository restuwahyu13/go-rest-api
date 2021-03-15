package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl login",
	})
}
