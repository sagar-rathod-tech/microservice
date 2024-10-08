package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment struct {
	ID      string  `json:"id"`
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

var payments = map[string]Payment{
	"1": {ID: "1", OrderID: "1", Amount: 100.50},
}

func main() {
	r := gin.Default()

	// CREATE
	r.POST("/payment", func(c *gin.Context) {
		var newPayment Payment
		if err := c.ShouldBindJSON(&newPayment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		payments[newPayment.ID] = newPayment
		c.JSON(http.StatusCreated, newPayment)
	})

	// READ
	r.GET("/payment/:id", func(c *gin.Context) {
		id := c.Param("id")
		payment, exists := payments[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
		c.JSON(http.StatusOK, payment)
	})

	// UPDATE
	r.PUT("/payment/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedPayment Payment
		if err := c.ShouldBindJSON(&updatedPayment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		payments[id] = updatedPayment
		c.JSON(http.StatusOK, updatedPayment)
	})

	// DELETE
	r.DELETE("/payment/:id", func(c *gin.Context) {
		id := c.Param("id")
		delete(payments, id)
		c.JSON(http.StatusOK, gin.H{"message": "Payment deleted"})
	})

	r.Run(":8083") // Run on port 8083
}
