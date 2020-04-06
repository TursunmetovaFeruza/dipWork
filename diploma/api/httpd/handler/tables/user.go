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

	} else if usertype == "master" {

	}
	// user := model_user.SignIn{}

	// var usertype string
	user := model_student.StudentGet{}
	res := db.QueryRow(query, userid)
	err := res.Scan(&user.Username, &user.Password, &user.Name, &user.Surname, &user.Fathername, &user.EnrollmentYear, &user.Course, &user.Language,
		&user.Group, &user.Email, &user.LectionId, &user.Faculty, &user.School, &user.Speciality)
	if err != nil {
		fmt.Println("signIn problem:", err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &user,
		})
	}

	// if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(user.Password)); err != nil {
	// 	c.JSON(200, gin.H{
	// 		"message": "Not correct",
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"message":  "Succes",
	// 		"user":     &user.Username,
	// 		"user_id":  &stored.UserInfo,
	// 		"userType": &usertype,
	// 	})
	// }
	defer db.Close()

}
