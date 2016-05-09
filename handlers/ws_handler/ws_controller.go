package wshandler

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gkiryaziev/go-ws-server/utils"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSController struct
type WSController struct {
	hub *Hub
}

// NewWsController return new WSController object.
func NewWsController(hub *Hub) *WSController {
	return &WSController{hub}
}

// WsHandler websocket handler.
func (wsc *WSController) WsHandler(w http.ResponseWriter, r *http.Request) {
	// get incoming connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// generate new id with size as md5 string
	uid := utils.GenerateRandomMD5String(16)

	// add id and connection to pool
	conn := NewConnection(ws, uid, wsc.hub)
	conn.hub.register <- conn

	// onClose
	defer func() {
		// delete id and connection from pool
		conn.hub.unregister <- conn
	}()

	go conn.writer()
	conn.reader()
}
