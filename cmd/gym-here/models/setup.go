package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbSSLMode := os.Getenv("DB_SSLMODE")

	// Construct the connection URL
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", DbHost, DbPort, DbUser, DbName, DbPassword, DbSSLMode)

	DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	if err := DB.AutoMigrate(&User{}, &Client{}, &Training{}, &Instructor{}).Error; err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}
}
