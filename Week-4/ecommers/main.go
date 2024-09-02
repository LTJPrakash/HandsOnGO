package main

import (
	"ecommers/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Server starting...")

	r := router.Router() // Initialize the router

	// Define the route for the home page
	r.HandleFunc("/", serveHome).Methods("GET")

	// CORS middleware setup
	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Adjust this to your frontend URL
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Wrap the router with CORS middleware
	handler := handlers.CORS(corsObj, corsHeaders, corsMethods)(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":5000", handler))
	fmt.Println("Listening at port 5000")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home for API in GOLANG 5000</h1>"))
}
