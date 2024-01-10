package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arifinhabibi/campus-go/app/controller"
	"github.com/gorilla/mux"
)

func Router() {
	// Initialize the router
    router := mux.NewRouter()

    // // Add routes
    router.HandleFunc("/items", controller.GetItems).Methods("GET")
    // router.HandleFunc("/items/{id}", GetItem).Methods("GET")
    // router.HandleFunc("/items", CreateItem).Methods("POST")
    // router.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
    // router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}