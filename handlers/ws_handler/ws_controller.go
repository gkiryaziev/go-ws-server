package ws_handler

import (
	"net/http"

	"../../utils"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsController struct {
	hub *hub
}

// ==========================
// Websocket Controller
// ==========================
func NewWsController(hub *hub) *wsController {
	return &wsController{hub}
}

// ==========================
// Websocket Handler
// ==========================
func (this *wsController) WsHandler(w http.ResponseWriter, r *http.Request) {
	// get incoming connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// generate new id with size as md5 string
	uid := utils.GenerateRandomMD5String(16)

	// add id and connection to pool
	conn := NewConnection(ws, uid, this.hub)
	conn.hub.register <- conn

	// onClose
	defer func() {
		// delete id and connection from pool
		conn.hub.unregister <- conn
	}()

	go conn.writer()
	conn.reader()
}
