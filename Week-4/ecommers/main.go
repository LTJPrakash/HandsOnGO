package main

import (
	"context"
	"ecommers/router" // Ensure this includes your OAuth utils
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "ecommers/docs" // Import the generated Swagger docs

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger" // Swagger handler for Gorilla Mux
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauth2Config *oauth2.Config
	oauthState   = generateStateCookie() // Use a dynamic state parameter
)

// @title E-Commerce API in GO
// @version 1.0
// @description This will be used for testing of API endpoints.
// @host localhost:5000
// @BasePath /
// @schemes http
func main() {
	fmt.Println("Server starting...")

	r := router.Router() // Initialize the router

	// Set up Google OAuth2 configuration
	oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	// Swagger route setup
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Define the route for the home page
	r.HandleFunc("/", serveHome).Methods("GET")

	// Google OAuth routes
	r.HandleFunc("/auth/google/login", handleGoogleLogin).Methods("GET")
	r.HandleFunc("/auth/google/callback", handleGoogleCallback).Methods("GET")

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

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauthState {
		http.Error(w, "State mismatch", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauth2Config.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer userInfo.Body.Close()

	data, err := io.ReadAll(userInfo.Body)
	if err != nil {
		http.Error(w, "Failed to read user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// generateStateCookie generates a random state value for OAuth2 security
func generateStateCookie() string {
	return "pseudo-random" // Replace with a more secure random value
}
