package utils

// ResponseMessage struct
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Info    string `json:"info"`
}

// ErrorMessage return message as json string
func ErrorMessage(status int, msg string) string {
	msgFinal := &ResponseMessage{status, msg, "/docs/api/errors"}
	result, _ := NewResultTransformer(msgFinal).ToJSON()
	return result
}
