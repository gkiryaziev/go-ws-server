package subscriberhandler

import "time"

// SubscriberModel struct
type SubscriberModel struct {
	ID       int64     `db:"id" json:"id"`
	DateTime time.Time `db:"datetime" json:"datetime"`
	UID      string    `db:"uid" json:"uid"`
	Topic    string    `db:"topic" json:"topic"`
}
