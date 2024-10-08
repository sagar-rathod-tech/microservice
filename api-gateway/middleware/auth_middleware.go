package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			// If no token, respond with Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Prevents the handler from continuing
			return
		}

		// Example token validation (in production, validate the token properly)
		if token != "Bearer my-secret-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Allow the request to proceed
		c.Next()
	}
}
