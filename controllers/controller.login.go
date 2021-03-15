package controller

import (
	"net/http"

	"github.com/gin-gionic/gin"
)

func LoginController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl",
	})
}
