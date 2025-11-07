package main

import (
	"go.mod/config"
	"go.mod/internal/routes"
	"go.mod/models"
	"log"
	"os"
)

func main() {
	config.Init()
	if err := config.DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	r := routes.SetupRouter()
	port := os.Getenv("APP_PORT")
	log.Printf("Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
