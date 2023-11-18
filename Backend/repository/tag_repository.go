package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
)

type TagRepository struct {
	DB *sql.DB
}

func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{DB: db}
}

func (tr *TagRepository) CreateTag(tag models.Tag) error {
	query := "INSERT INTO Tags (TagName) VALUES (?)"
	_, err := tr.DB.Exec(query, tag.TagName)
	if err != nil {
		log.Printf("Error creating tag: %v", err)
		return fmt.Errorf("failed to create tag: %v", err)
	}

	return nil
}

func (tr *TagRepository) GetTags(UserID) ([]models.Tag, error) {
	query := "SELECT * FROM Tags"
	rows, err := tr.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching tags: %v", err)
		return nil, fmt.Errorf("failed to fetch tags: %v", err)
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.TagID, &tag.TagName); err != nil {
			log.Printf("Error scanning tag rows: %v", err)
			return nil, fmt.Errorf("failed to scan tag rows: %v", err)
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
