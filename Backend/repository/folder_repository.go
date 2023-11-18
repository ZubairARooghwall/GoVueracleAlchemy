package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
)

type FolderRepository struct {
	DB *sql.DB
}

func NewFolderRepository(db *sql.DB) *FolderRepository {
	return &FolderRepository{DB: db}
}

func (fr *FolderRepository) CreateFolder(folder models.Folder) error {
	query := "INSERT INTO Folders (FolderName, UserID, CreationTime) VALUES (?, ?, CURRENT_TIMESTAMP)"
	_, err := fr.DB.Exec(query, folder.FolderName, folder.User.UserID)
	if err != nil {
		log.Printf("Error creating folder %v", err)
		return err
	}

	return nil
}

func (fr *FolderRepository) GetFolderByID(folderID int) (*models.Folder, error) {
	query := "SELECT * FROM Folders WHERE FolderID = ?"
	row := fr.DB.QueryRow(query, folderID)

	var folder models.Folder
	err := row.Scan(&folder.FolderID, &folder.FolderName, &folder.User.UserID, &folder.CreationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("folder with ID %d not found", folderID)
		}
		log.Printf("Error retrieving folder: %v", err)
		return nil, err
	}

	return &folder, nil
}

func (fr *FolderRepository) GetFoldersByUserID(UserID int) (*[]models.Folder, error) {
	query := "SELECT * FROM Folders WHERE UserID = ?"
	rows, err := fr.DB.Query(query, UserID)
	if err != nil {
		log.Printf("Error retrieving folders for user: %v", err)
		return nil, err
	}
	rows.Close()

	var folders []models.Folder
	for rows.Next() {
		var folder models.Folder
		err = rows.Scan(&folder.FolderID, &folder.FolderName, &folder.User.UserID, &folder.CreationDate)
		if err != nil {
			log.Printf("Error scanning folder rows: %v", err)
			return nil, err
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over folder rows: %v", err)
		return nil, err
	}

	return &folders, nil
}

func (fr *FileRepository) DeleteFolder(folderID int) error {
	query := "DELETE FROM Folders WHERE FolderID = ?"
	_, err := fr.DB.Exec(query, folderID)
	if err != nil {
		log.Printf("Error deleting folder: %v", err)
		return err
	}

	return nil
}
