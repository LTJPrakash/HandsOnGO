package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the system
// @Description User model for registration and login
// @Name User
// @Property id string "User ID" example("605c72ef4f1f4f4f4f4f4f4f")
// @Property username string "Username" example("john_doe")
// @Property password string "Password" example("securepassword")
// @Property email string "Email" example("john.doe@example.com")
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
}
