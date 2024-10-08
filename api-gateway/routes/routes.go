package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-tech/api-gateway/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// Grouping routes for user service
	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.AuthenticationMiddleware()) // Authentication only for user service
		userRoutes.Any("/*path", func(c *gin.Context) {
			forwardRequest(c, "http://localhost:8081"+c.Request.URL.Path) // Forwarding request correctly
		})
	}

	// Routes for order service, applying logging middleware only
	orderRoutes := r.Group("/order")
	{
		orderRoutes.Use(middleware.LoggingMiddleware()) // Logging only for order service
		orderRoutes.Any("/*path", func(c *gin.Context) {
			forwardRequest(c, "http://localhost:8082"+c.Request.URL.Path) // Forwarding request correctly
		})
	}

	// Payment service, no middleware applied
	r.POST("/payment", func(c *gin.Context) { // Changed to POST without trailing slash
		forwardRequest(c, "http://localhost:8083/payment")
	})

	// If you want to allow other methods like GET or DELETE, you can add:
	r.Any("/payment/*path", func(c *gin.Context) { // Allow all methods for payment path
		forwardRequest(c, "http://localhost:8083/payment"+c.Param("path")) // Correct path handling
	})
}

// Helper function to forward requests
func forwardRequest(c *gin.Context, url string) {
	fullURL := url + "?" + c.Request.URL.RawQuery // Append query parameters

	// Create a new request to forward
	req, err := http.NewRequest(c.Request.Method, fullURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
