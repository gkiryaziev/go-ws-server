package subscriber_handler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

type subscriberController struct {
	service *subscriberService
}

// NewSubscriberController return new subscriberController object.
func NewSubscriberController(db *sqlx.DB) *subscriberController {
	return &subscriberController{newSubscriberService(db)}
}

// GetSubscribers return all subscribers.
func (sc *subscriberController) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := sc.service.getSubscribers().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
