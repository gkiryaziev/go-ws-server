package topichandler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

// TopicService struct
type TopicService struct {
	db *sqlx.DB
}

// newTopicService return new TopicService object.
func newTopicService(db *sqlx.DB) *TopicService {
	return &TopicService{db}
}

// getTopics return all topics.
func (ts *TopicService) getTopics() *utils.ResultTransformer {

	topics := []TopicModel{}

	err := ts.db.Select(&topics, "select * from topics order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{Status: "ok", Count: len(topics), Data: topics}
	result := utils.NewResultTransformer(header)

	return result
}
