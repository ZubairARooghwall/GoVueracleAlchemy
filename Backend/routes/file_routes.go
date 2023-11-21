package routes

import (
	"github.com/ZubairARooghwall/GoVueracleAlchemy/controllers"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/middleware"
	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(router *gin.Engine) {
	fileRoutes := router.Group("/files")
	fileRoutes.Use(middleware.AuthMiddleware)

	fileRoutes.POST("/upload", controllers.Uploadfile)
	fileRoutes.GET("/", controllers.GetAllFiles)
	fileRoutes.GET("/:id", controllers.GetFilesByID)
	fileRoutes.DELETE("/:id", controllers.DeleteFileByID)
}
