package main

import (
	"github.com/arifinhabibi/campus-go/app"
	"github.com/arifinhabibi/campus-go/app/database"
)

func main() {
    // database connection
    database.Connection()

    // router
    app.Router();
}
