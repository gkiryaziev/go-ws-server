package topic_handler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

type topicService struct {
	db *sqlx.DB
}

// newTopicService return new topicService object.
func newTopicService(db *sqlx.DB) *topicService {
	return &topicService{db}
}

// getTopics return all topics.
func (ts *topicService) getTopics() *utils.ResultTransformer {

	topics := []TopicModel{}

	err := ts.db.Select(&topics, "select * from topics order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(topics), topics}
	result := utils.NewResultTransformer(header)

	return result
}
