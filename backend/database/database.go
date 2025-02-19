package database

import (
	"fmt"
	"os"

	"github.com/SabdaSankalla/ecommerce-pharmacy/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{})

	DB = db
}
