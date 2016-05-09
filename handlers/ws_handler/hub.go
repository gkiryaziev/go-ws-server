package wshandler

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

//Hub struct
type Hub struct {
	connections map[string]*Connection
	broadcast   chan *Broadcast
	register    chan *Connection
	unregister  chan *Connection
	service     *wsService
}

// NewHub return new Hub object.
func NewHub(db *sqlx.DB) *Hub {
	return &Hub{
		connections: make(map[string]*Connection),
		broadcast:   make(chan *Broadcast),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		service:     newWSService(db),
	}
}

// Run Hub's main method.
func (h *Hub) Run() {
	for {
		select {
		// register new connection
		case conn := <-h.register:
			h.connections[conn.uid] = conn

			// get remote address
			wsRAddress := conn.ws.RemoteAddr().String()

			// add log
			h.service.addLog(conn.uid, wsRAddress, "Connected")

		// unregister connection
		case conn := <-h.unregister:
			if _, ok := h.connections[conn.uid]; ok {
				close(conn.send)

				delete(h.connections, conn.uid)

				// unsubscribe from all topic
				h.service.unSubscribeAll(conn.uid)

				// get remote address
				wsRAddress := conn.ws.RemoteAddr().String()

				// add log
				h.service.addLog(conn.uid, wsRAddress, "Disconnected")
			}

		// read incoming message
		case b := <-h.broadcast:
			// add log
			h.service.addLog(b.uid, b.address, string(b.message))

			// unmarshal message
			var msg WSMessage
			err := json.Unmarshal(b.message, &msg)
			if err != nil {
				break
			}

			switch msg.Action {
			case "SUBSCRIBE":
				h.service.subscribe(msg.Topic, b.uid)
			case "UNSUBSCRIBE":
				h.service.unSubscribe(msg.Topic, b.uid)
			case "PUBLISH":
				// get all subscribers by topic name
				subscribers := h.service.getSubscribers(msg.Topic)
				// check if subscribers are greater then zero
				if len(subscribers) > 0 {
					// get subscriber from list
					for _, subscriberID := range subscribers {
						// check if subscriber is not me
						if subscriberID != b.uid {
							// get subscriber connection by id
							if conn, ok := h.connections[subscriberID]; ok {
								select {
								// send message
								case conn.send <- b.message:
								default:
									close(conn.send)
									delete(h.connections, conn.uid)
								}
							}
						}
					}
				}
			}
		}
	}
}
