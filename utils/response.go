package utils

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	StatusCode int
	Method     string
	Message    string
	Data       interface{}
}

func APIResponse(c *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	c.JSON(StatusCode, jsonResponse)
}
