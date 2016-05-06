package subscriber_handler

import "time"

type SubscriberModel struct {
	Id       int64     `db:"id" json:"id"`
	DateTime time.Time `db:"datetime" json:"datetime"`
	Uid      string    `db:"uid" json:"uid"`
	Topic    string    `db:"topic" json:"topic"`
}
