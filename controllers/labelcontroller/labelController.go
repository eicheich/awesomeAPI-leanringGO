package labelcontroller

import (
	// "awesomeAPI/middleware"
	"awesomeAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func ShowLabel(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	// Find label
	label, err := models.GetLabelsByUserID(int(userID.(uint)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": label})
}
func CreateLabel(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var input models.Label
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	label := models.Label{Name: input.Name, Color: input.Color, UserID: int(userID.(uint))}
	err := models.DB.Create(&label).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": label})
}
