package subscriber_handler

import (
	"ws.server/models"
	"ws.server/utils"

	"github.com/jmoiron/sqlx"
)

type subscriberService struct {
	db *sqlx.DB
}

func newSubscriberService(db *sqlx.DB) *subscriberService {
	return &subscriberService{db}
}

// get subscribers.
// note: get topic as name instead of id
func (this *subscriberService) getSubscribers() *utils.ResultTransformer {

	subscribers := []SubscriberModel{}

	err := this.db.Select(&subscribers,
		"select id, datetime, uid, (select name from topics where id = subscribers.topic_id) as topic from subscribers")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(subscribers), subscribers}
	result := utils.NewResultTransformer(header)

	return result
}