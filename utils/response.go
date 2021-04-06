package utils

type Responses struct {
	StatusCode uint8
	Method     string
	Type       string
	Message    string
	Data       interface{}
}

func APIResponse(Message string, StatusCode uint8, Method, Type string, Data interface{}) Responses {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Type:       Type,
		Message:    Message,
		Data:       Data,
	}
	return jsonResponse
}
