package models

import "time"

type UserSession struct {
	SessionID     int       `json:"user_session_id"`
	User          User      `json:"user"`
	SessionToken  string    `json:"session_token"`
	ExpiryDate    time.Time `json:"expiry_date"`
	UserIPAddress string    `json:"user_ip_address"`
	Browser       string    `json:"browser"`
}
