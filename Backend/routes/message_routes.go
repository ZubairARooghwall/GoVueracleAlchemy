package routes

import (
	"github.com/ZubairARooghwall/GoVueracleAlchemy/controllers"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/middleware"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

func SetUpMessageRoutes(router *gin.Engine, messageRepo *repository.MessageRepository) {
	messageController := controllers.MessageController{messageRepo: messageRepo}

	messageRoutes := router.Group("/messages")
	{
		messageRoutes.Use(middleware.AuthMiddleware)

		messageRoutes.POST("/", messageController.CreateMessage)
		messageRoutes.GET("/:receiverID", messageController.GetMessagesBetweenUsers)
	}

}
