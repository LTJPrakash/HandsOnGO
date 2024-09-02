package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user document in the MongoDB collection
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
}
