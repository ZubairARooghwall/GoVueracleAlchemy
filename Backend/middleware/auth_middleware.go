package middleware

import (
	"net/http"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userSessionRepo *repository.UserSessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken := c.GetHeader("Authorization")

		if sessionToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, err := userSessionRepo.ValidateSessionToken(sessionToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
