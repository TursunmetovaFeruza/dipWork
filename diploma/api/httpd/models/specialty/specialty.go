package model

//Speciality
type Speciality struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	ShortName string `json:"shortname" bson:"shortname"`
	FacultyId int    `json:"facultyId" bson:"faculty_id"`
	SchoolId  int    `json:"schoolId" bson:"school_id"`
}

//items of Speciality
type Specialities []Speciality
