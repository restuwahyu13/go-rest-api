package util

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) Response {
	var jsonResponse Response
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
