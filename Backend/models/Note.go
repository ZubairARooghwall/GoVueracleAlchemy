package models

import "time"

type Note struct {
	NoteID       int       `json:"node_id"`
	Meeting      Meeting   `json:"meeting"`
	Owner        User      `json:"owner"`
	CreationTime time.Time `json:"creation_time"`
	Content      string    `json:"content"`
	Permission   rune      `json:"permission"`
}
