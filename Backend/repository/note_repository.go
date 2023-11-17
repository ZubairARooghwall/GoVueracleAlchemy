package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
)

type NoteRepository struct {
	DB *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{DB: db}
}

func (nr *NoteRepository) CreateNote(note models.Note) error {
	query := "INSERT INTO Notes (UserID, Content, CreationTime) VALUES(?, ?, CURRENT_TIMESTAMP)"
	_, err := nr.DB.Exec(query, note.Owner, note.Content)
	if err != nil {
		log.Printf("Error creating note: %v", err)
		return fmt.Errorf("failed to create note: %v", err)
	}

	return nil
}

func (nr *NoteRepository) GetNoteByUserID(userID int) ([]models.Note, error) {
	query := "SELECT * FROM Notes WHERE UserID = ?"
	rows, err := nr.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching notes: %v", err)
		return nil, fmt.Errorf("failed to fetch notes: %v", err)
	}

	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.NoteID, &note.Owner, &note.Content); err != nil {
			log.Printf("Error scanning note rows: %v", err)
			return nil, fmt.Errorf("failed to scan note rows: %v", err)
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (nr *NoteRepository) DeleteNote(noteID int) error {
	query := "DELETE FROM Notes WHERE NoteID = ?"
	_, err := nr.DB.Exec(query, noteID)
	if err != nil {
		log.Printf("Error deleting note: %v", err)
		return fmt.Errorf("failed to delete note: %v", err)
	}

	return nil
}
