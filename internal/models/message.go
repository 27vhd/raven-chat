package models

import "time"

type Message struct {
	ID        int
	Content   string
	Username  string
	Timestamp time.Time
}
