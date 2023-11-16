package models

import "time"

type File struct {
	FileID       int       `json:"file_id"`
	FileName     string    `json:"file_name"`
	UniqueName   string    `json:"unique_name"`
	FileSize     int       `json:"file_size"`
	CreationTime time.Time `json:"creation_time"`
	Owner        User      `json:"owner"`
	Folder       Folder    `json:"folder"`
	Tag          Tag       `json:"tag"`
}
