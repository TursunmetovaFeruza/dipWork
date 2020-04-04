package model

//Language
type Language struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	ShortName string `json:"shortname" bson:"shortname"`
}

//items of Language
type Languages []Language
