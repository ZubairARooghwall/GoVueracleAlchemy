package models

import "time"

type Folder struct {
	FolderID     int       `json:"folder_id"`
	User         User      `json:"user"`
	FolderName   string    `json:"folder_name"`
	CreationDate time.Time `json:"creation_date"`
	FolderSize   int       `json:"folder_size"`
}
