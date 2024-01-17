package controller

import (
	"encoding/json"
	"net/http"

	"github.com/arifinhabibi/campus-go/app/models/users"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func Login(w  http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ e", "application/json")
    users := users.Shows()
    json.NewEncoder(w).Encode(users)
}