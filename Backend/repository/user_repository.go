package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
	"golang.org/x/crypto/bcrypt"
)

func generateSaltValue() (string, error) {
	saltBytes := make([]byte, 64)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(saltBytes), nil
}

func hashPasswordWithSalt(password, salt string) (string, error) {
	saltBytes, err := base64.URLEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	combined := append([]byte(password), saltBytes...)

	hashedPassword, err := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (ur *UserRepository) ValidateUserCredentials(email, password string) (int, error) {
	query := "SELECT UserID, Password, Salt FROM Users WHERE Email = ?"
	row := ur.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserID, &user.Password, &user.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}

		log.Printf("Error retrieving user credentials: %v", err)
		return 0, fmt.Errorf("failed to retrieve user credenetials: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+string(user.Salt))); err != nil {
		return 0, errors.New("invalid password")
	}

	return user.UserID, nil
}

func (ur *UserRepository) storeSessionInformation(userID int, sessionToken string, expiryDate time.Time) error {
	query := "INSERT INTO UserSessions (UserID, SessionToken, ExpiryDate) VALUES (?, ?, ?)"
	_, err := ur.DB.Exec(query, userID, sessionToken, expiryDate)
	if err != nil {
		log.Printf("Error storing session information: %v", err)
		return err
	}

	return nil
}

func (ur *UserRepository) GenerateSessionToken(userID int) (string, error) {
	sessionToken, err := generateUniqueSessionToken()
	if err != nil {
		return "", fmt.Errorf("failed to generate unique session token: %v", err)
	}

	expiryDate := time.Now().Add(3 * 30 * 24 * time.Hour)

	err = ur.storeSessionInformation(userID, sessionToken, expiryDate)
	if err != nil {
		return "", fmt.Errorf("failed to store session information: %v", err)
	}

}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	salt, err := generateSaltValue()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %v", err)
	}

	hashedPassword, err := hashPasswordWithSalt(user.Password, salt)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	query := "INSERT INTO Users (Username, Email, Password, Salt, Role, Education, CreationTime) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)"

	_, err = ur.DB.Exec(query, user.Username, user.Email, hashedPassword, salt, user.Education)
	if err != nil {
		// Log or handle the error appropriately
		log.Printf("Error creating user: %v", err)
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := "SELECT UserID, Username, Email, Password, Salt, Role, Education, CreationTime FROM Users WHERE UserID = ?"
	row := ur.DB.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Salt, &user.UserRole, &user.Education, &user.CreationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", userID)
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT UserID, Username, Email, Password, Salt, Role, Education, CreationTime FROM Users WHERE Username = ?"
	row := ur.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Salt, &user.UserRole, &user.Education, &user.CreationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT UserID, Username, Email, Password, Salt, Role, Education, CreationTime FROM Users WHERE Email = ?"
	row := ur.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Salt, &user.UserRole, &user.Education, &user.CreationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) DeleteUser(userID int) error {
	query := "DELETE FROM Users WHERE UserID = :1"
	_, err := ur.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

func (ur *UserRepository) ListAllUsers() ([]models.User, error) {
	query := "SELECT * FROM Users"
	rows, err := ur.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Salt, &user.UserRole, &user.Education, &user.CreationTime)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user rows: %v", err)
		}

		users = append(users, user)
	}

	return users, nil
}
