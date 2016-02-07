package utils

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Info    string `json:"info"`
}

// ========================
// return message as json string
// ========================
func ErrorMessage(status int, msg string) string {
	msg_final := &ResponseMessage{status, msg, "/docs/api/errors"}
	result, _ := NewResultTransformer(msg_final).ToJson()
	return result
}
