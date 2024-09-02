package router

import (
	"ecommers/controller"

	"github.com/gorilla/mux"
)

// Router initializes and returns a new router with all the API routes defined
func Router() *mux.Router {
	router := mux.NewRouter() // Create a new router instance

	// Define the route for getting all products
	router.HandleFunc("/api/products", controller.GetProducts).Methods("GET")

	// Define the route for creating a new product
	router.HandleFunc("/api/product", controller.CreateProduct).Methods("POST")

	// Define the route for updating a product
	router.HandleFunc("/api/product/{id}", controller.UpdateProduct).Methods("PUT")

	// Define the route for deleting a specific product
	router.HandleFunc("/api/product/{id}", controller.DeleteProduct).Methods("DELETE")

	// Define the route for getting a single product by ID
	router.HandleFunc("/api/product/{id}", controller.GetProductByID).Methods("GET")

	// Define the route for user registration
	router.HandleFunc("/api/register", controller.Register).Methods("POST")

	// Define the route for user login
	router.HandleFunc("/api/login", controller.Login).Methods("POST")

	return router // Return the configured router
}
