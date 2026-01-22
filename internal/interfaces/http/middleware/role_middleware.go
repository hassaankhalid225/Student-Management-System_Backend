package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/sms-backend/internal/domain/entities"
)

func RoleMiddleware(roles ...entities.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user role not found"})
			c.Abort()
			return
		}

		role := userRole.(entities.UserRole)
		authorized := false
		for _, r := range roles {
			if r == role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
