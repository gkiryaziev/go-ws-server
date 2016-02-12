package ws_handler

import (
	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn
	uid  string
	send chan []byte
	hub  *Hub
}

type broadcast struct {
	uid     string
	address string
	message []byte
}

// Connection constructor.
func NewConnection(ws *websocket.Conn, uid string, hub *Hub) *connection {
	return &connection{
		ws:   ws,
		uid:  uid,
		send: make(chan []byte, 256),
		hub:  hub,
	}
}

// Connection reader.
func (this *connection) reader() {
	b := &broadcast{}
	for {
		// read incoming message
		_, message, err := this.ws.ReadMessage()
		if err != nil {
			break
		}

		b.uid = this.uid
		b.address = this.ws.RemoteAddr().String()
		b.message = message

		this.hub.broadcast <- b
	}
	this.ws.Close()
}

// Connection writer.
func (this *connection) writer() {
	for message := range this.send {
		err := this.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	this.ws.Close()
}
