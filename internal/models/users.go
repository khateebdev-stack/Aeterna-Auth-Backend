package models

import "time"
// This is  User struct represents a user in our database.
type User struct{
	ID string `json:"id"`
	Email string  `json:"email"`
	HashedPassword string `json:"_"` // We won't serialize this field to JSON
CreatedAt time.Time `json:"created_at"`
}


// RegisterRequest structure for user Registration
type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}