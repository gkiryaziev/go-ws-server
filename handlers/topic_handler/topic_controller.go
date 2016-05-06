package topic_handler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

type topicController struct {
	service *topicService
}

// NewTopicController return new topicController object.
func NewTopicController(db *sqlx.DB) *topicController {
	return &topicController{newTopicService(db)}
}

// GetTopics return all topics.
func (tc *topicController) GetTopics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	topics, err := tc.service.getTopics().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, topics)
}
