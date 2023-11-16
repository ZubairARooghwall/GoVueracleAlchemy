package models

import "time"

type Meeting struct {
	MeetingID    int       `json:"meeting_id"`
	MeetingTitle string    `json:"meeting_title"`
	Description  string    `json:"description"`
	Organizer    User      `json:"organizer"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}
