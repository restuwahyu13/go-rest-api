package util

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func Strigify(payload map[string]interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) Response {
	var jsonResponse Response
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Error(err.Error())
	}

	return jsonResponse
}
