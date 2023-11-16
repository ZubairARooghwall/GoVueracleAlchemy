package models

import "time"

type Tag struct {
	TagName      string    `json:"name"`
	CreationTime time.Time `json:"creation_time"`
	Color        string    `json:"color"`
}
