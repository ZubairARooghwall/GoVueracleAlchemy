package models

import "time"

type Permission struct {
	PermissionID   int       `json:"permission_id"`
	File           File      `json:"file"`
	Sender         User      `json:"sender"`
	Receiver       User      `json:"receiver"`
	PermissionType string    `json:"permission_type"`
	CreationTime   time.Time `json:"creation_time"`
}
