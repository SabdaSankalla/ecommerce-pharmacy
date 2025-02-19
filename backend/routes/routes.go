package routes

import (
	"github.com/SabdaSankalla/ecommerce-pharmacy/handlers"
	"github.com/SabdaSankalla/ecommerce-pharmacy/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)
	router.GET("/products", handlers.GetProducts)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/products", handlers.CreateProduct)
		protected.PUT("/products/:id", handlers.UpdateProduct)
		protected.DELETE("/products/:id", handlers.DeleteProduct)
	}
}
