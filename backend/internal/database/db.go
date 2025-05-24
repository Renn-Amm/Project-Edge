package database

import (
	"log"
	"os"

	"github.com/Renn-Amm/Project-Edge/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	DB = db
	return DB
}
