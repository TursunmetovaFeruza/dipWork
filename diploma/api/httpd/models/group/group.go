package model

//Group
type Group struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	FacultyId    int    `json:"facultyId" bson:"faculty_id"`
	SchoolId     int    `json:"schoolId" bson:"school_id"`
	SpecialityId int    `json:"specialityId" bson:"speciality_id"`
}

//items of Group
type Groups []Group
