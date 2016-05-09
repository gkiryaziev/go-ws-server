package topichandler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/utils"
)

// TopicController struct
type TopicController struct {
	service *TopicService
}

// NewTopicController return new TopicController object.
func NewTopicController(db *sqlx.DB) *TopicController {
	return &TopicController{newTopicService(db)}
}

// GetTopics return all topics.
func (tc *TopicController) GetTopics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	topics, err := tc.service.getTopics().ToJSON()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, topics)
}
