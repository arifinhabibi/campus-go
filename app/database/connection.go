package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var url string

func Connection(){
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

	db, err := sql.Open("mysql", url)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
	
	err = db.Ping()
    if err != nil {
		panic(err)
    }
	fmt.Println("Database success Connection!")
}