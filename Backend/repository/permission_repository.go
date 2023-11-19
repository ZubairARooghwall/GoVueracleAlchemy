package repository

import (
	"GoVueracleAlchemy/models"
	"database/sql"
	"fmt"
	"log"
)

type PermissionRepository struct {
	DB *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (pr *PermissionRepository) CreatePermission(permission models.Permission) error {
	query := "INSERT INTO Permissions (Sender, UserID, FileID, PermissionType, CreationTime)"
	_, err := pr.DB.Exec(query, permission.Sender.UserID, permission.Receiver.UserID, permission.File.FileID, string(permission.PermissionType))
	if err != nil {
		log.Printf("Error creating permission: %v", err)
		return fmt.Errorf("failed to create permission: %v", err)
	}

	return nil
}

func (pr *PermissionRepository) GetPermissionByUserID(userID int) ([]models.Permission, error) {
	query := "SELECT * FROM Permissions WHERE UserID = ? ORDER BY CreationTime"
	rows, err := pr.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching permissions: %v", err)
		return nil, fmt.Errorf("failed to fetch permissions: %v", err)
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		var permission models.Permission
		if err := rows.Scan(&permission.PermissionID, &permission.Sender.UserID, &permission.File.FileID, &permission.PermissionType, &permission.CreationTime); err != nil {
			log.Printf("Error scanning permission rows: %v", err)
			return nil, fmt.Errorf("failed to scan permission rows: %v", err)
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}
