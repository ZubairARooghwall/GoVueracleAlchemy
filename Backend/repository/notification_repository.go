package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (nr *NotificationRepository) GetNotificationByUserID(userID int) ([]models.Notification, error) {
	query := "SELECT * FROM Notifications WHERE ToUser = ? ORDER BY CreationTime"
	rows, err := nr.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error fetching notification: %v", err)
		return nil, fmt.Errorf("failed to fetch notifications: %v", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(&notification.NotificationID, &notification.ToUser.UserID, &notification.NotificationType, &notification.CreationTime, &notification.FromUser.UserID, &notification.Content, &notification.IsRead); err != nil {
			log.Printf("Error scanning notification rows: %v", err)
			return nil, fmt.Errorf("failed to scan notification rows: %v", err)
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (nr *NotificationRepository) MarkNotificationAsRead(notificationID int) error {
	query := "UPDATE Notifications SET IsRead = 'Y' WHERE NotificationID = ?"
	_, err := nr.DB.Exec(query, notificationID)
	if err != nil {
		log.Printf("Error making notification as read: %v", err)
		return fmt.Errorf("failed to make notification as read: %v", err)
	}

	return nil
}
