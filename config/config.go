package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB        *gorm.DB
	SecretKey string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	dsn := fmt.Sprintf(
		"host =%s user=%s password=@%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	DB = db
	SecretKey = os.Getenv("SECRET_KEY")
	log.Println("Successfully connected to database")
}
