package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a default gin router
	r := gin.Default()

	// Define a route handler for the root path
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Framework!",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
} 