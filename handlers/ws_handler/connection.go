package wshandler

import (
	"github.com/gorilla/websocket"
)

// Connection struct
type Connection struct {
	ws   *websocket.Conn
	uid  string
	send chan []byte
	hub  *Hub
}

// Broadcast struct
type Broadcast struct {
	uid     string
	address string
	message []byte
}

// NewConnection return new Connection object.
func NewConnection(ws *websocket.Conn, uid string, hub *Hub) *Connection {
	return &Connection{
		ws:   ws,
		uid:  uid,
		send: make(chan []byte, 256),
		hub:  hub,
	}
}

// reader is Connection reader.
func (c *Connection) reader() {
	b := &Broadcast{}
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

// writer is Connection writer.
func (c *Connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
