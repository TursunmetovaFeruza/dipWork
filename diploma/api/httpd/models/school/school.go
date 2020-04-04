package model

//School
type School struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	FacultyId int    `json:"facultyId" bson:"faculty_id"`
	ShortName string `json:"shortname" bson:"shortname"`
}

//items of School
type Schools []School
