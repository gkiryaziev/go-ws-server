package topic_handler

import "time"

type TopicModel struct {
	Id       int64     `db:"id" json:"id"`
	DateTime time.Time `db:"datetime" json:"datetime"`
	Name     string    `db:"name" json:"name"`
}
