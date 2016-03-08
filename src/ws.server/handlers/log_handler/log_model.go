package log_handler

import "time"

type LogModel struct {
	Id            int64     `db:"id" json:"id"`
	DateTime      time.Time `db:"datetime" json:"datetime"`
	Uid           string    `db:"uid" json:"uid"`
	RemoteAddress string    `db:"remote_address" json:"remote_address"`
	Message       string    `db:"message" json:"message"`
}
