package handlers

import (
	"net/http"
	"time"

	"github.com/SabdaSankalla/ecommerce-pharmacy/database"
	"github.com/SabdaSankalla/ecommerce-pharmacy/models"

	"github.com/gin-gonic/gin"
)

// Get All Products
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// Create Product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// Update Product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.UpdatedAt = time.Now()

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Delete Product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
