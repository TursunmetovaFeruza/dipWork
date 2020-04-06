package model

//Speciality
type Student struct {
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

type StudentGet struct {
	Username       string      `json:"username" bson:"username"`
	Password       string      `json:"password" bson:"password"`
	Name           string      `json:"name" bson:"name"`
	Surname        string      `json:"surname" bson:"surname"`
	Fathername     string      `json:"fathername" bson:"fathername"`
	EnrollmentYear int         `json:"enrollmentyear" bson:"enrollment_year"`
	Course         int         `json:"course" bson:"course"`
	Language       string      `json:"lang" bson:"language"`
	Group          string      `json:"group" bson:"group"`
	Email          string      `json:"email" bson:"email"`
	LectionId      interface{} `json:"lectionId" bson:"lection_id"`
	Faculty        string      `json:"faculty" bson:"faculty"`
	School         string      `json:"school" bson:"school"`
	Speciality     string      `json:"speciality" bson:"speciality"`
}

//items of Speciality
type Students []Student
