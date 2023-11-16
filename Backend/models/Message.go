package models

import "time"

type Message struct {
	MessageID    int       `json:"message_id"`
	Sender       User      `json:"sender"`
	Receiver     User      `json:"receiver"`
	IsSent       rune      `json:"is_sent"`
	IsRead       rune      `json:"is_read"`
	CreationTime time.Time `json:"creation_time"`
	Content      string    `json:"content"`
}
