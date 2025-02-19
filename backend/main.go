package main

import (
	"github.com/SabdaSankalla/ecommerce-pharmacy/database"
	"github.com/SabdaSankalla/ecommerce-pharmacy/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Initialize Router
	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r)

	// Run Server
	r.Run(":5050")
}
