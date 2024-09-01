package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Netflix represents a movie document in the MongoDB collection
type Netflix struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`         // Unique identifier for the movie (MongoDB ObjectID)
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`     // Title of the movie
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"` // Indicates whether the movie has been watched
	Genre   string             `json:"genre,omitempty" bson:"genre,omitempty"`     // New field: Genre of the movie
}
