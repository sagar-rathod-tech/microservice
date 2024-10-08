package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// After request is processed
		latency := time.Since(start)
		log.Printf("%s %s | Latency: %v | Status: %d",
			c.Request.Method, c.Request.URL.Path, latency, c.Writer.Status())
	}
}
