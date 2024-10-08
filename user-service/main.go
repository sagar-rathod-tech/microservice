package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{
	"1": {ID: "1", Name: "John Doe", Age: 30},
}

func main() {
	r := gin.Default()

	// CREATE
	r.POST("/user", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users[newUser.ID] = newUser
		c.JSON(http.StatusCreated, newUser)
	})

	// READ
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, exists := users[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// UPDATE
	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users[id] = updatedUser
		c.JSON(http.StatusOK, updatedUser)
	})

	// DELETE
	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		delete(users, id)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	})

	r.Run(":8081") // Run on port 8081
}
