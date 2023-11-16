package models

import "time"

type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Salt         []byte    `json:"salt"`
	UserRole     string    `json:"user_role"`
	Education    string    `json:"education"`
	CreationTime time.Time `json:"creation_time"`
}
