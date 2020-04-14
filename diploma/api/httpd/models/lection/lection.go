package model

//Lection
type Lection struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

//items of Lection
type Lections []Lection
