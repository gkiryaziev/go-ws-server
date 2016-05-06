package subscriber_handler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

type subscriberService struct {
	db *sqlx.DB
}

// newSubscriberService return new subscriberService object.
func newSubscriberService(db *sqlx.DB) *subscriberService {
	return &subscriberService{db}
}

// getSubscribers return all subscribers.
// note: get topic as name instead of id
func (ss *subscriberService) getSubscribers() *utils.ResultTransformer {

	subscribers := []SubscriberModel{}

	err := ss.db.Select(&subscribers,
		"select id, datetime, uid, (select name from topics where id = subscribers.topic_id) as topic from subscribers")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(subscribers), subscribers}
	result := utils.NewResultTransformer(header)

	return result
}
