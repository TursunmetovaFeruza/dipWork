package model

//Faculty
type Faculty struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	ShortName string `json:"shortname" bson:"shortname"`
}

//items of Faculty
type Faculties []Faculty
