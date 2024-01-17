package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var url string

func Connection() (*gorm.DB) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	  // Access environment variables
	  dbHost := os.Getenv("DB_HOST")
	  dbPort := os.Getenv("DB_PORT")
	  dbUser := os.Getenv("DB_USER")
	  dbPassword := os.Getenv("DB_PASSWORD")
	  dbName := os.Getenv("DB_NAME")


	if dbPassword == "" {
		url = dbUser+"@tcp("+ dbHost +":"+ dbPort +")/"+dbName
	} else {
		url = dbUser+":"+ dbPassword +"@tcp("+ dbHost +":"+ dbPort +")/"+dbName
	}

	// db, err := sql.Open("mysql", url)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
    if err != nil {
        panic(err.Error())
    }

	return db
}