package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
)

type UserSessionRepository struct {
	DB *sql.DB
}

func NewUserSessionRepository(db *sql.DB) *UserSessionRepository {
	return &UserSessionRepository{DB: db}
}

func (usr *UserSessionRepository) GenerateSessionToken(userID int) (string, error) {
	sessionToken, err := generateUniqueSessionToken()
	if err != nil {
		return "", fmt.Errorf("failed to generate session token: %v", err)
	}

	expiryDate := time.Now().Add(3 * 30 * 24 * time.Hour)

	err = usr.storeSessionInformation(userID, sessionToken, expiryDate)
	if err != nil {
		return "", fmt.Errorf("failed to store session information: %v", err)
	}

	return sessionToken, nil
}

func (usr *UserSessionRepository) storeSessionInformation(userID int, sessionToken string, expiryDate time.Time) error {
	query := "INSERT INTO UserSessions (UserID, SessionToken, ExpiryDate) VALUES (?, ?, ?)"
	_, err := usr.DB.Exec(query, userID, sessionToken, expiryDate.UTC())
	if err != nil {
		log.Printf("Error storing session information: %v", err)
		return err
	}

	return nil
}

func (usr *UserSessionRepository) GetActiveSessionByUserID(userID int) ([]models.UserSession, error) {
	query := "SELECT * FROM UserSessions WHERE UserID = ? AND ExpiryDate > CURRENT_TIMESTAMP"
	rows, err := usr.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching active sessions: %v", err)
		return nil, err
	}

	defer rows.Close()

	var sessions []models.UserSession
	for rows.Next() {
		var session models.UserSession
		if err := rows.Scan(&session.SessionID, &session.User.UserID, &session.SessionToken, &session.ExpiryDate, &session.UserIPAddress, &session.Browser); err != nil {
			log.Printf("Error scanning session rows: %v", err)
			return nil, err
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (usr *UserSessionRepository) GetSessionByID(sessionID int) (*models.UserSession, error) {
	query := "SELECT * FROM UserSessions WHERE SessionID = ?"
	row := usr.DB.QueryRow(query, sessionID)

	var userSession models.UserSession
	err := row.Scan(&userSession.SessionID, &userSession.User.UserID, &userSession.SessionToken, &userSession.ExpiryDate, &userSession.UserIPAddress, &userSession.Browser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("session with ID %d not found", sessionID)
		}
		log.Printf("Error retrieving session information: %v", err)
		return nil, err
	}

	return &userSession, nil
}

func (usr *UserSessionRepository) DeleteSession(sessionID int) error {
	query := "DELETE FROM UserSessions WHERE SessionID = ?"
	_, err := usr.DB.Exec(query, sessionID)
	if err != nil {
		log.Printf("Error deleting session: %v", err)
		return err
	}

	return nil
}

func generateUniqueSessionToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	sessionToken := base64.URLEncoding.EncodeToString(randomBytes)

	return sessionToken, nil
}
