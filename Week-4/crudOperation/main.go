package main

import (
	"crud/router" // Import the router package from the 'crud' module
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("DB API") // Print a message to indicate that the DB API server is starting

	// Initialize the router from the router package
	r := router.Router() // Ensure this function name matches the one in your 'router.go'

	// Define the route for the home page
	r.HandleFunc("/", serveHome).Methods("GET")

	fmt.Println("Server starting...") // Print a message indicating the server is starting
	// Start the server on port 5000 and log any fatal errors
	log.Fatal(http.ListenAndServe(":5000", r))

	fmt.Println("Listening at port 5000") // Print a message indicating the server is listening on port 5000
}

// serveHome handles requests to the home page
func serveHome(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to text/html for the response
	// w.Header().Set("Content-Type", "text/html")

	// Write a simple HTML message as the response body
	w.Write([]byte("<h1>Home for API in GOLANG 5000</h1>"))
}
