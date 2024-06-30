package labelcontroller

import (
	// "awesomeAPI/middleware"
	// "awesomeAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func ShowLabel(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": userID})
}