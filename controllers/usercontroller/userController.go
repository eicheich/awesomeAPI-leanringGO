package usercontroller

import (
	"awesomeAPI/models"
	"net/http"
	"github.com/gin-gonic/gin"

)


func Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Create(c *gin.Context) {
	var input models.CreateUserInput
	// check if username is unique
	username := input.Username
	var user models.User
	models.DB.Where("username = ?", username).First(&user)
	if user.Username == username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// check if email is unique
	email := input.Email
	models.DB.Where("email = ?", email).First(&user)
	if user.Email == email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = models.User{Username: input.Username, Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Show(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Update(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username is unique
	username := input.Username
	var existingUser models.User
	models.DB.Where("username = ?", username).First(&existingUser)
	if existingUser.Username == username && existingUser.ID != user.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Check if email is unique
	email := input.Email
	models.DB.Where("email = ?", email).First(&existingUser)
	if existingUser.Email == email && existingUser.ID != user.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Delete(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}