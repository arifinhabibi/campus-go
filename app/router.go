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
    router.HandleFunc("/user", controller.Login).Methods("POST")

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}