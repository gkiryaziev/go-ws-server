package topic_handler

import (
	"ws.server/models"
	"ws.server/utils"

	"github.com/jmoiron/sqlx"
)

type topicService struct {
	db *sqlx.DB
}

func newTopicService(db *sqlx.DB) *topicService {
	return &topicService{db}
}

// get topics.
func (this *topicService) getTopics() *utils.ResultTransformer {

	topics := []TopicModel{}

	err := this.db.Select(&topics, "select * from topics order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(topics), topics}
	result := utils.NewResultTransformer(header)

	return result
}
