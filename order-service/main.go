package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

var orders = map[string]Order{
	"1": {ID: "1", UserID: "1", Amount: 100.50},
}

func main() {
	r := gin.Default()

	// CREATE
	r.POST("/order", func(c *gin.Context) {
		var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orders[newOrder.ID] = newOrder
		c.JSON(http.StatusCreated, newOrder)
	})

	// READ
	r.GET("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		order, exists := orders[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	// UPDATE
	r.PUT("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedOrder Order
		if err := c.ShouldBindJSON(&updatedOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orders[id] = updatedOrder
		c.JSON(http.StatusOK, updatedOrder)
	})

	// DELETE
	r.DELETE("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		delete(orders, id)
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
	})

	r.Run(":8082") // Run on port 8082
}
