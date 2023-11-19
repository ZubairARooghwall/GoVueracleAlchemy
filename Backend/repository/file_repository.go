package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileRepository struct {
	DB       *sql.DB
	BasePath string
}

func NewFileRepository(db *sql.DB, basePath string) *FileRepository {
	return &FileRepository{DB: db, BasePath: basePath}
}

func (fr *FileRepository) SaveFile(fileBytes []byte, fileName string, userID int) error {
	filePath := filepath.Join(fr.BasePath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(fileBytes)
	if err != nil {
		return err
	}

	query := "INSERT INTO Files (FileName, FilePath, UserID, UploadTime) VALUES (?, ?, ?, ?)"
	_, err = fr.DB.Exec(query, fileName, filePath, userID, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileRepository) GetFilePath(fileName string) string {
	return filepath.Join(fr.BasePath, fileName)
}

func (fr *FileRepository) DeleteFile(fileName string) error {
	filePath := filepath.Join(fr.BasePath, fileName)

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	query := "DELETE FROM Files WHERE FileName = ?"
	_, err = fr.DB.Exec(query, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileRepository) GetAllFilesByUserID(userID int) ([]models.File, error) {
	query := "SELECT * FROM Files WHERE Owner = ?"
	rows, err := fr.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching files: %v", err)
		return nil, fmt.Errorf("failed to fetch files: %v", err)
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.FileID, &file.FileName, &file.Owner.UserID, &file.CreationTime); err != nil {
			log.Printf("Error scanning file rows: %v", err)
			return nil, fmt.Errorf("failed to scan file rows: %v", err)
		}

		files = append(files, file)
	}

	return files, nil
}
