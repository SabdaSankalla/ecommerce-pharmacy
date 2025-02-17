package main

import (
	"github.com/gin-gonic/gin"
	"github.com/SabdaSankalla/ecommerce-pharmacy/backend/database"
    "github.com/SabdaSankalla/ecommerce-pharmacy/backend/routes"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Initialize Router
	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r)

	// Run Server
	r.Run(":8080")
}
