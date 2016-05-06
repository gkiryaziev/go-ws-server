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

// NewConnection return new connection object.
func NewConnection(ws *websocket.Conn, uid string, hub *Hub) *connection {
	return &connection{
		ws:   ws,
		uid:  uid,
		send: make(chan []byte, 256),
		hub:  hub,
	}
}

// reader is connection reader.
func (c *connection) reader() {
	b := &broadcast{}
	for {
		// read incoming message
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		b.uid = c.uid
		b.address = c.ws.RemoteAddr().String()
		b.message = message

		c.hub.broadcast <- b
	}
	c.ws.Close()
}

// writer is connection writer.
func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
