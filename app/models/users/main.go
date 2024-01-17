package users

import (
	"github.com/arifinhabibi/campus-go/app/database"
)

type User struct {
    ID       int
    Username string
    Email    string
	Password string
}

func Register() {
	dbConnection := database.Connection()
	
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(50) NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(100) NULL
	);`

	dbConnection.Exec(createTable)
}

func Shows()([]User){
	dbConnection := database.Connection()

	// selectQuery := `SELECT * FROM users`
	
	var users []User

	if err := dbConnection.Find(&users).Error; err != nil {
        return nil
    }
	
	return users
}

func Create(payload User) {
	dbConnection := database.Connection()

	dbConnection.Create(payload)
}
