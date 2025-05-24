package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Renn-Amm/Project-Edge/internal/handlers"
	"github.com/Renn-Amm/Project-Edge/internal/repositories"
	"github.com/Renn-Amm/Project-Edge/internal/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set in the environment")
	}

	userRepo := repositories.NewUserRepo()
	userService := services.NewUserService(userRepo)

	r := gin.Default()
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/api/v1/login", userHandler.Login)
	r.POST("/api/v1/signup", userHandler.SignUp)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
