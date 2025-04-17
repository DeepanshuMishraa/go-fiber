package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DeepanshuMishraa/go-fiber/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	DB *gorm.DB
}

var Database dbInstance

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Calcutta", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	log.Println("Database connection established")

	// migrations
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations.....")

	db.AutoMigrate(&models.User{}, &models.Question{})

	Database.DB = db
}
