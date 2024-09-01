package controller

import (
	"context"
	"crud/model" // Import the model package from the 'crud' module
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" // Import the Gorilla Mux package for routing
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection settings
const dbName = "netflix"
const collectionName = "watchlist"

// MongoDB collection instance
var collection *mongo.Collection

// Initialize MongoDB connection
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var connectionString = os.Getenv("MONGO_URL") // Fetch MongoDB connection string from environment variable

	// Create a new MongoDB client option with the connection string
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("MongoDB connection success")

	// Set the collection instance
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

// Insert a new movie into the collection and return its ID
func insertOneMovie(movie model.Netflix) primitive.ObjectID {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Failed to insert movie:", err)
	}
	fmt.Println("Inserted 1 movie id:", inserted.InsertedID)
	return inserted.InsertedID.(primitive.ObjectID)
}

// Update the 'watched' status of a movie
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("Failed to update movie:", err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

// Delete a specific movie by its ID
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Failed to delete movie:", err)
	}
	fmt.Println("Movie deleted with count:", deleteCount)
}

// Delete all movies in the collection
func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal("Failed to delete all movies:", err)
	}
	fmt.Println("Deleted movie number:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Retrieve all movies from the collection
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal("Failed to find movies:", err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal("Failed to decode movie:", err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Controller function to handle GET requests for all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// Controller function to handle POST requests for creating a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")     // Set the content type to JSON
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Correct CORS header

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Insert the movie and get the inserted ID
	insertedID := insertOneMovie(movie)
	movie.Id = insertedID
	json.NewEncoder(w).Encode(movie)
}

// Controller function to handle PUT requests for marking a movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")    // Set the content type to JSON
	w.Header().Set("Access-Control-Allow-Methods", "PUT") // Correct CORS header

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Controller function to handle DELETE requests for deleting a specific movie
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // Set the content type to JSON
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // Correct CORS header

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Controller function to handle DELETE requests for deleting all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // Set the content type to JSON
	w.Header().Set("Access-Control-Allow-Methods", "DELETE") // Correct CORS header

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
