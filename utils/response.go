package utils

type Responses struct {
	StatusCode uint8
	Method     string
	Message    string
	Data       interface{}
}

func APIResponse(Message string, StatusCode uint8, Method string, Data interface{}) Responses {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}
	return jsonResponse
}
