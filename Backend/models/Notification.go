package models

import "time"

type Notification struct {
	NotificationID   int       `json:"notification_id"`
	ToUser           User      `json:"to_user"`
	NotificationType string    `json:"notification_type"`
	CreationTime     time.Time `json:"creation_time"`
	FromUser         User      `json:"from_user"`
	Content          string    `json:"content"`
	IsRead           rune      `json:"is_read"`
}
