package subscriberhandler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

// SubscriberController struct
type SubscriberController struct {
	service *SubscriberService
}

// NewSubscriberController return new SubscriberController object.
func NewSubscriberController(db *sqlx.DB) *SubscriberController {
	return &SubscriberController{newSubscriberService(db)}
}

// GetSubscribers return all subscribers.
func (sc *SubscriberController) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := sc.service.getSubscribers().ToJSON()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
