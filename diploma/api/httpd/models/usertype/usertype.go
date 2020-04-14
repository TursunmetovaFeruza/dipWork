package model

//UserType
type UserType struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

//items of UserType
type UserTypes []UserType
