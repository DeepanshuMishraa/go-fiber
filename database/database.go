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

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Connected to database")

	//Todo: add migrations
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	//Reasons for {} -> it is a pointer to the struct
	// and we are passing the address of the struct to the AutoMigrate function

	Database.DB = db
}
