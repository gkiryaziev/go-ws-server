package wshandler

// WSMessage struct
type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Data   string `json:"data"`
}
