package database

import (
	"fmt"
	"log"
	"os"
   
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
   )
   
   var DB *gorm.DB
   
   func DatabaseInit() {
   
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Access environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")

	var err error
   
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
   
	if err != nil {
	 log.Fatal("Failed to Connect to database...")
	}
   
	fmt.Println("Connecting to database...")
}