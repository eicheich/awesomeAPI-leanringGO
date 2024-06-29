package labelcontroller

import (
	"awesomeAPI/models"
	"net/http"
	"github.com/gin-gonic/gin"

)

// show label by user id
func ShowLabel(c *gin.Context) {
	var labels []models.Label
	if err := models.DB.Where("user_id = ?", c.Param("id")).Find(&labels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Label not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": labels})
}