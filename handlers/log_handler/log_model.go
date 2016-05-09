package loghandler

import "time"

// LogModel struct
type LogModel struct {
	ID            int64     `db:"id" json:"id"`
	DateTime      time.Time `db:"datetime" json:"datetime"`
	UID           string    `db:"uid" json:"uid"`
	RemoteAddress string    `db:"remote_address" json:"remote_address"`
	Message       string    `db:"message" json:"message"`
}
