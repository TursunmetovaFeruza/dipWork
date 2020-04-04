package model

//Speciality
type Speciality struct {
	ID             int    `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Surname        string `json:"surname" bson:"surname"`
	Fathername     string `json:"fathername" bson:"fathername"`
	EnrollmentYear int    `json:"enrollmentyear" bson:"enrollment_year"`
	CourseId       int    `json:"courseId" bson:"curs_id"`
	LanguageId     int    `json:"langId" bson:"lang_id"`
	GroupId        int    `json:"groupId" bson:"group_id"`
	Email          string `json:"email" bson:"email"`
	LectionId      []int  `json:"lectionId" bson:"lection_id"`
	FacultyId      int    `json:"facultyId" bson:"faculty_id"`
	SchoolId       int    `json:"schoolId" bson:"school_id"`
	SpecialityId   int    `json:"specialityId" bson:"speciality_id"`
}

//items of Speciality
type Specialities []Speciality
