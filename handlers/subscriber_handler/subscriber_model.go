package subscriber_model

import "time"

type SubscriberModel struct {
	Id       int64     `db:"id" json:"id"`
	DateTime time.Time `db:"datetime" json:"datetime"`
	Uid      string    `db:"uid" json:"uid"`
	TopicId  int64     `db:"topic_id" json:"topic_id"`
}
