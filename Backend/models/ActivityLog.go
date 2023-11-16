package models

import "time"

type ActivityLog struct {
	ActivityID   int       `json:"activity_id"`
	Activity     string    `json:"activity"`
	CreationTime time.Time `json:"creation_time"`
	User         User      `json:"user"`
	File         File      `json:"file"`
}
