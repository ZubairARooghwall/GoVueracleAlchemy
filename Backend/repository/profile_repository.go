package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (pr *ProfileRepository) GetProfileByID(profileID int) (*models.Profile, error) {
	query := "SELECT * FROM Profiles WHERE ProfileID = ?"
	row := pr.DB.QueryRow(query, profileID)

	var profile models.Profile
	err := row.Scan(&profile.ProfileID, &profile.User, &profile.ProfilePicture, &profile.Status, &profile.Biography)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("profile with ID %d not found", profileID)
		}
		log.Printf("Error retrieving profile: %v", err)
		return nil, err
	}

	return &profile, nil
}

func (pr *ProfileRepository) CreateProfile(profile models.Profile) error {
	query := "INSERT INTO Profiles (ProfileID, UserID, ProfilePicture, Status, Biography) VALUES (?, ?, ?, ?, ?)"
	_, err := pr.DB.Exec(query, profile.ProfileID, profile.User, profile.ProfilePicture, profile.Status, profile.Biography)
	if err != nil {
		log.Printf("Error creating profile: %v", err)
		return err
	}

	return nil
}

func (pr *ProfileRepository) DeleteProfile(profileID int) error {
	query := "DELETE FROM Profiles WHERE ProfileID = ?"
	_, err := pr.DB.Exec(query, profileID)
	if err != nil {
		log.Printf("Error deleting profile: %v", err)
	}

	return nil
}
