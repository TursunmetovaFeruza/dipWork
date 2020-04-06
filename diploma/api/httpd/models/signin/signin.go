package model

import (
	"time"
)

//User
type SignIn struct {
	ID        int       `json:"id" bson:"id"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UserInfo  int       `json:"userinfoId" bson:"id"`
}

//Users
type SignIns []SignIn
