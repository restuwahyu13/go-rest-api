package util

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int
	Method     string
	Message    string
	Data       interface{}
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}
	ctx.JSON(StatusCode, jsonResponse)
}
