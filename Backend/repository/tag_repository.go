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
	query := "INSERT INTO Tags (TagName, CreationTime, Color) VALUES (?, CURRENT_TIMESTAMP, ?)"
	_, err := tr.DB.Exec(query, tag.TagName, tag.Color)
	if err != nil {
		log.Printf("Error creating tag: %v", err)
		return fmt.Errorf("failed to create tag: %v", err)
	}

	return nil
}

func (tr *TagRepository) GetTagByName(tagName string) (*models.Tag, error) {
	query := "SELECT * FROM Tags WHERE TagName = ?"
	row := tr.DB.QueryRow(query, tagName)

	var tag models.Tag
	err := row.Scan(&tag.TagName, &tag.CreationTime, &tag.Color)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag with name %s not found", tagName)
		}
		log.Printf("Error retrieving tag: %v", err)
		return nil, err
	}

	return &tag, nil
}

func (tr *TagRepository) GetTags() ([]models.Tag, error) {
	query := "SELECT * FROM Tags ORDER BY CreationTime"
	rows, err := tr.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching tags: %v", err)
		return nil, fmt.Errorf("failed to fetch tags: %v", err)
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.TagName, &tag.CreationTime, &tag.Color); err != nil {
			log.Printf("Error scanning tag rows: %v", err)
			return nil, fmt.Errorf("failed to scan tag rows: %v", err)
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (tr *TagRepository) DeleteTag(tagName string) error {
	query := "DELETE FROM Tags WHERE TagName = ?"
	_, err := tr.DB.Exec(query, tagName)
	if err != nil {
		log.Printf("Error deleting tag: %v", err)
		return fmt.Errorf("failed to deleted tag: %v", err)
	}

	return nil
}
