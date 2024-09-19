package controller

import (
	"context"
	"ecommers/model"
	"ecommers/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection    *mongo.Collection
	productCollection *mongo.Collection
)

// run the first in file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve MongoDB connection string from environment variable
	connectionString := os.Getenv("MONGO_URI")
	if connectionString == "" {
		log.Fatal("MONGO_URI is not set in the environment")
	}

	// Create a MongoDB client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the MongoDB server to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")

	// Initialize collections
	userCollection = client.Database("ecommerce").Collection("users")
	productCollection = client.Database("ecommerce").Collection("products")
}

// Register godoc
// @Summary Register a new user
// @Description Registers a new user with username, password, and email, and hashes the password
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "User data with username, password, and email"
// @Success 201 {object} model.User "User successfully created"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error creating user"
// @Router /api/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login godoc
// @Summary Login a user and return a JWT token
// @Description Authenticates a user and returns a JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "Invalid username or password"
// @Failure 500 {string} string "Error generating token"
// @Router /api/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	var foundUser model.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(user.Password, foundUser.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
