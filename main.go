package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func main() {
    // Initialize the router
    router := mux.NewRouter()

    // // Add routes
    // router.HandleFunc("/items", GetItems).Methods("GET")
    // router.HandleFunc("/items/{id}", GetItem).Methods("GET")
    // router.HandleFunc("/items", CreateItem).Methods("POST")
    // router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
    // router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

    // Start the server
    port := 8080
    fmt.Printf("Server is running on port %d...\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
