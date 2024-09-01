package router

import (
	"crud/controller" // Import the controller package from the 'crud' module

	"github.com/gorilla/mux" // Import the Gorilla Mux router package
)

// Router initializes and returns a new router with all the API routes defined
func Router() *mux.Router {
	router := mux.NewRouter() // Create a new router instance

	// Define the route for getting all movies
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")

	// Define the route for creating a new movie
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	// Define the route for marking a movie as watched
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	// Define the route for deleting a specific movie
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	// Define the route for deleting all movies
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router // Return the configured router
}
