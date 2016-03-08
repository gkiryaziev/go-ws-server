package subscriber_handler

import (
	"fmt"
	"net/http"

	"ws.server/utils"

	"github.com/jmoiron/sqlx"
)

type subscriberController struct {
	service *subscriberService
}

func NewSubscriberController(db *sqlx.DB) *subscriberController {
	return &subscriberController{newSubscriberService(db)}
}

// get subscribers.
func (this *subscriberController) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	logs, err := this.service.getSubscribers().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, logs)
}
