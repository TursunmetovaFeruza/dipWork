package tables

import (
	"fmt"
	model_student "newsfeeder/httpd/models/student"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetUser(c *gin.Context) {
	db := PostSQLConfig()
	var query string
	userid := c.Request.URL.Query().Get("id")
	usertype := c.Request.URL.Query().Get("type")
	if usertype == "student" {

		query = "select username, password, students.name as name, students.surname as surname , students.fathername as fathername, " +
			"students.enrollment_year as enrollment_year, curses_number.number as course, language.name as language, groups.name as group, " +
			"email, lection_id, faculty.name as faculty, school.name as school, specialty.name as speciality " +
			"FROM users " +
			"INNER JOIN students ON users.user_info_id =students.id " +
			"INNER JOIN curses_number ON students.curs_id = curses_number.id  " +
			"INNER JOIN language ON students.lang_id = language.id  " +
			"INNER JOIN groups ON students.group_id = groups.id " +
			"INNER JOIN faculty ON students.faculty_id = faculty.id " +
			"INNER JOIN school ON students.school_id = school.id " +
			"INNER JOIN specialty ON students.specialty_id = specialty.id " +
			"where users.id =$1"
	} else if usertype == "admin" {
		query = "select username, password " +
			"FROM users " +
			"where users.id =$1"
	} else if usertype == "master" {
		query = "select username, password, master.name as name, master.surname as surname , master.fathername as fathername, " +
			"email, lection_id, faculty.name as faculty, school.name as school " +
			"FROM users " +
			"INNER JOIN master ON users.user_info_id = master.id " +
			"INNER JOIN faculty ON master.faculty_id = faculty.id " +
			"INNER JOIN school ON master.school_id = school.id " +
			"where users.id =$1"
	}
	// user := model_user.SignIn{}

	// var usertype string
	res := db.QueryRow(query, userid)
	user := model_student.StudentGet{}
	var err interface{}
	if usertype == "student" {
		err = res.Scan(&user.Username, &user.Password, &user.Name, &user.Surname, &user.Fathername, &user.EnrollmentYear, &user.Course, &user.Language,
			&user.Group, &user.Email, &user.LectionId, &user.Faculty, &user.School, &user.Speciality)
	} else if usertype == "admin" {
		err = res.Scan(&user.Username, &user.Password)
	} else if usertype == "master" {
		err = res.Scan(&user.Username, &user.Password, &user.Name, &user.Surname, &user.Fathername, &user.Email, &user.LectionId,
			&user.Faculty, &user.School)
	}

	if err != nil {
		fmt.Println("signIn problem:", err)
	} else {
		fmt.Println(user)
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &user,
		})
	}
	defer db.Close()
}

// select   students.name as name, students.surname as surname , students.fathername as fathername,
// students.enrollment_year as enrollment_year, curses_number.number as course, language.name as language, groups.name as group,
//  email, faculty.name as faculty, school.name as school, specialty.name as speciality
//   FROM students

//  INNER JOIN curses_number ON students.curs_id = curses_number.id
//  INNER JOIN language ON students.lang_id = language.id
//  INNER JOIN groups ON students.group_id = groups.id
//  INNER JOIN faculty ON students.faculty_id = faculty.id
// INNER JOIN school ON students.school_id = school.id
// INNER JOIN specialty ON students.specialty_id = specialty.id

func GetAllStudents(c *gin.Context) {
	db := PostSQLConfig()
	query := "select   students.name as name, students.surname as surname , students.fathername as fathername, " +
		"students.enrollment_year as enrollment_year, curses_number.number as course, language.name as language, groups.name as group, " +
		"email, lection_id, faculty.name as faculty, school.name as school, specialty.name as speciality " +
		"FROM students " +
		"INNER JOIN curses_number ON students.curs_id = curses_number.id  " +
		"INNER JOIN language ON students.lang_id = language.id  " +
		"INNER JOIN groups ON students.group_id = groups.id " +
		"INNER JOIN faculty ON students.faculty_id = faculty.id " +
		"INNER JOIN school ON students.school_id = school.id " +
		"INNER JOIN specialty ON students.specialty_id = specialty.id "
	users := []model_student.StudentGet{}
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	for rows.Next() {
		user := model_student.StudentGet{}
		err := rows.Scan(&user.Name, &user.Surname, &user.Fathername, &user.EnrollmentYear, &user.Course, &user.Language,
			&user.Group, &user.Email, &user.LectionId, &user.Faculty, &user.School, &user.Speciality)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &users,
		})
	}
	defer db.Close()
}

// select   master.name as name, master.surname as surname , master.fathername as fathername,
//  email, faculty.name as faculty, school.name as school
//   FROM master
//  INNER JOIN faculty ON master.faculty_id = faculty.id
// INNER JOIN school ON master.school_id = school.id

func GetAllMaters(c *gin.Context) {
	db := PostSQLConfig()
	query := "select   master.name as name, master.surname as surname , master.fathername as fathername, " +
		"email, lection_id, faculty.name as faculty, school.name as school " +
		"FROM master " +
		"INNER JOIN faculty ON master.faculty_id = faculty.id " +
		"INNER JOIN school ON master.school_id = school.id "
	users := []model_student.StudentGet{}
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	for rows.Next() {
		user := model_student.StudentGet{}
		err := rows.Scan(&user.Name, &user.Surname, &user.Fathername,
			&user.Email, &user.LectionId, &user.Faculty, &user.School)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &users,
		})
	}
	defer db.Close()
}
