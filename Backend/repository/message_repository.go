package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (mr *MessageRepository) CreateMessage(message models.Message) error {
	query := "INSERT INTO Messages (Sender, Receiver, Content, IsSent, IsRead, CreationTime) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP)"

	_, err := mr.DB.Exec(query, message.Sender.UserID, message.Receiver.UserID, message.Content, message.IsSent, message.IsRead)
	if err != nil {
		log.Printf("Error creating message: %v", err)
		return fmt.Errorf("failed to create message: %v", err)
	}
	return nil
}

func (mr *MessageRepository) GetMessageBetweenUsers(senderID, receiverID int) ([]models.Message, error) {
	query := "SELECT * FROM Messages WHERE (SenderID = ? AND ReceiverID = ?) OR (SenderID = ? AND ReceiverID = ?) ORDER BY CreationTime"
	rows, err := mr.DB.Query(query, senderID, receiverID, receiverID, senderID)
	if err != nil {
		log.Printf("Error fetching message: %v", err)
		return nil, fmt.Errorf("failed to fetch mesage: %v", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.MessageID, &message.Sender.UserID, &message.Receiver.UserID, &message.Content, &message.IsSent, &message.IsRead, &message.CreationTime); err != nil {
			log.Printf("Error scanning message rows: %v", err)
			return nil, fmt.Errorf("failed to scan message rows: %v", err)
		}

		messages = append(messages, message)
	}

	return messages, nil
}
