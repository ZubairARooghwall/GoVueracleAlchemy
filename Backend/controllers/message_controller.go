package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	MessageRepo *repository.MessageRepository
}

func NewMessageController(messageRepo *repository.MessageRepository) *MessageController {
	return &MessageController{MessageRepo: messageRepo}
}

func (mc *MessageController) CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.MessageRepo.CreateMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (mc *MessageController) GetMessagesBetweenUsers(c *gin.Context) {
	senderID, err := strconv.Atoi(c.Param("senderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender ID"})
		return
	}

	receiverID, err := strconv.Atoi(c.Param("receiverID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	messages, err := mc.MessageRepo.GetMessageBetweenUsers(senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
