package model

//Finger
type Finger struct {
	ID        int    `json:"id" bson:"id"`
	Finger    []byte `json:"finger" bson:"finger"`
	StudentId int    `json:"fstudentid" bson:"fstudentid"`
}

//items of Finger
type Fingers []Finger
