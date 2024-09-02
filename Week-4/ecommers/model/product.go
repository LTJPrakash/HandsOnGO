package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product represents a product document in the MongoDB collection
type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
}
