package models

type Friendship struct {
	FriendshipID int    `json:"friendship_id"`
	User         User   `json:"sender"`
	Friend       User   `json:"receiver"`
	Status       string `json:"status"`
}
