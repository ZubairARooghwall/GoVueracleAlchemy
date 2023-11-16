package repository

import (
	"../models"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
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

	hashedPassword, err := crypto.bcrypt.GenerateFromPassword(combined, crypto.bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO Users (Username, Email, Password, Salt, Role, Education, CreationTime) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)"

	salt, err := generateSaltValue()
	if err != nil {
		log.Printf("Error Failed to generate salt value: %v", err)
	}

	hashedPassword, err := hashPasswordWithSalt(user.Password, salt)

	_, err = ur.DB.Exec(query, user.Username, user.Email, hashedPassword, salt, user.Education)
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}

	return nil
}
