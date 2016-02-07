package topic_handler

import (
	"fmt"
	"net/http"

	"../../utils"

	"github.com/jmoiron/sqlx"
)

type topicController struct {
	service *topicService
}

func NewTopicController(db *sqlx.DB) *topicController {
	return &topicController{newTopicService(db)}
}

// ========================
// get topics
// ========================
func (this *topicController) GetTopics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	topics, err := this.service.getTopics().ToJson()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, utils.ErrorMessage(500, err.Error()))
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, topics)
}
