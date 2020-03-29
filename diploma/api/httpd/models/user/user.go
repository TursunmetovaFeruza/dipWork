package model

import (
	"time"
)

//User
type User struct {
	ID        int       `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Address   string    `json:"address" bson:"address"`
	Age       int       `json:"age" bson:"age"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

//Users
type Users []User
