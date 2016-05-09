package subscriberhandler

import (
	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-ws-server/models"
	"github.com/gkiryaziev/go-ws-server/utils"
)

// SubscriberService struct
type SubscriberService struct {
	db *sqlx.DB
}

// newSubscriberService return new SubscriberService object.
func newSubscriberService(db *sqlx.DB) *SubscriberService {
	return &SubscriberService{db}
}

// getSubscribers return all subscribers.
// note: get topic as name instead of id
func (ss *SubscriberService) getSubscribers() *utils.ResultTransformer {

	subscribers := []SubscriberModel{}

	err := ss.db.Select(&subscribers,
		"select id, datetime, uid, (select name from topics where id = subscribers.topic_id) as topic from subscribers")
	if err != nil {
		panic(err)
	}

	header := models.Header{Status: "ok", Count: len(subscribers), Data: subscribers}
	result := utils.NewResultTransformer(header)

	return result
}
