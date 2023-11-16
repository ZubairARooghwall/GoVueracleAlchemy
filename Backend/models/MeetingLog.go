package models

import "time"

type MeetingLog struct {
	MeetingLogID int       `json:"meeting_log_id"`
	Meeting      Meeting   `json:"meeting"`
	CreationTime time.Time `json:"creation_time"`
	User         User      `json:"user"`
	Activity     string    `json:"activity"`
	Description  string    `json:"description"`
}
