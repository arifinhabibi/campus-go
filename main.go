package main

import (
	"github.com/arifinhabibi/campus-go/app"
	"github.com/arifinhabibi/campus-go/app/database"
	"github.com/arifinhabibi/campus-go/app/models"
)

func main() {
    // database connection
    database.Connection()
    models.Register()

    // router
    app.Router();

}