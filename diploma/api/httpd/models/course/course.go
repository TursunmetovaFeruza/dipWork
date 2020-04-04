package model

//Course
type Course struct {
	ID     int `json:"id" bson:"id"`
	Number int `json:"number" bson:"number"`
}

//items of Course
type Courses []Course
