package repository

import (
	"database/sql"
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
