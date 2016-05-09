package topichandler

import "time"

// TopicModel struct
type TopicModel struct {
	ID       int64     `db:"id" json:"id"`
	DateTime time.Time `db:"datetime" json:"datetime"`
	Name     string    `db:"name" json:"name"`
}
