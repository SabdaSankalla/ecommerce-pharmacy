package handlers

import (
	"net/http"

	"github.com/SabdaSankalla/ecommerce-pharmacy/database"
	"github.com/SabdaSankalla/ecommerce-pharmacy/models"
	"github.com/gin-gonic/gin"
)

func RegisterUserTransaction(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func LoginUserTransaction(c *gin.Context) {
	// Belum Implementasi login dengan JWT
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
