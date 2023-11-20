package routes

import (
	"github.com/ZubairARooghwall/GoVueracleAlchemy/controllers"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/middleware"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(router *gin.Engine, userRepo *repository.UserRepository, userSessionRepo *repository.UserSessionRepository) {
	userController := controllers.UserController{UserRepo: *userRepo}

	userRoutes := router.Group("/users")
	{
		userRoutes.Use(middleware.AuthMiddleware)
		userRoutes.GET("/", userController.ListAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.DELETE("/delete/:id", userController.DeleteUserByID)

	}
}
