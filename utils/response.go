package util

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	Results    interface{} `json:"results"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(errValidator map[string]interface{}) *ErrorResponse {

	var errors ErrorResponse
	// var storeData = []string{}
	// var storeResult = append(storeData, Data)

	errDataCollection := make(map[string][]map[string]interface{})

	for i, v := range errValidator {
		errorResults := make(map[string]interface{})
		errorResults[i] = v
		errDataCollection["errors"] = append(errDataCollection["errors"], errorResults)
	}

	// errors.StatusCode = StatusCode
	// errors.Method = Method
	errors.Results = errDataCollection

	// fmt.Println(errDataCollection)

	return &errors

	// ctx.JSON(StatusCode, errors)
	// defer ctx.AbortWithStatus(StatusCode)
}
