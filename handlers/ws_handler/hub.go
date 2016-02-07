package ws_handler

import (
	"fmt"
	"time"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type hub struct {
	connections map[*connection]bool
	broadcast   chan *broadcast
	register    chan *connection
	unregister  chan *connection
	service     *wsService
}

// ==========================
// Hub constructor
// ==========================
func NewHub(db *sqlx.DB) *hub {
	return &hub{
		connections: make(map[*connection]bool),
		broadcast:   make(chan *broadcast),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		service:     newWSService(db),
	}
}

// ==========================
// Hub main method
// ==========================
func (this *hub) Run() {
	for {
		select {
		// register new connection
		case conn := <-this.register:
			this.connections[conn] = true

			// get remote address
			wsRAddress := conn.ws.RemoteAddr().String()

			// add log
			this.service.addLog(conn.uid, wsRAddress, "Connected")

		// unregister connection
		case conn := <-this.unregister:
			if _, ok := this.connections[conn]; ok {
				close(conn.send)
				delete(this.connections, conn)

				// unsubscribe from all topic
				this.service.unSubscribeAll(conn.uid)

				// get remote address
				wsRAddress := conn.ws.RemoteAddr().String()

				// add log
				this.service.addLog(conn.uid, wsRAddress, "Disconnected")
			}

		// read incoming message
		case b := <-this.broadcast:

			start := time.Now()

			// add log
			this.service.addLog(b.uid, b.address, string(b.message))

			// unmarshal message
			var msg WSMessage
			err := json.Unmarshal(b.message, &msg)
			if err != nil {
				break
			}

			switch msg.Action {
			case "SUBSCRIBE":
				this.service.subscribe(msg.Topic, b.uid)
			case "UNSUBSCRIBE":
				this.service.unSubscribe(msg.Topic, b.uid)
			case "PUBLISH":
				subscribers := this.service.getSubscribers(msg.Topic)
				if len(subscribers) > 0 {
					for _, subscriberId := range subscribers {
						for conn := range this.connections {
							if conn.uid == subscriberId {
								select {
								case conn.send <- b.message:
								default:
									close(conn.send)
									delete(this.connections, conn)
								}
							}
						}
					}
				}
				fmt.Println("PUBLISH elapsed", time.Since(start))
			}
		}
	}
}
