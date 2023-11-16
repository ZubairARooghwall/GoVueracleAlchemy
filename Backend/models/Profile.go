package models

type Profile struct {
	ProfileID      int    `json:"profile_id"`
	User           User   `json:"user"`
	ProfilePicture string `json:"profile_picture"`
	Status         rune   `json:"status"`
	Biography      string `json:"biography"`
}
