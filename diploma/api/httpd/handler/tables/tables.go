package tables

import (
	"fmt"
	model_course "newsfeeder/httpd/models/course"
	model_faculty "newsfeeder/httpd/models/faculty"
	model_groupe "newsfeeder/httpd/models/group"
	model_language "newsfeeder/httpd/models/language"
	model_lection "newsfeeder/httpd/models/lection"
	model_school "newsfeeder/httpd/models/school"
	model_specialty "newsfeeder/httpd/models/specialty"
	model_usertype "newsfeeder/httpd/models/usertype"

	"github.com/gin-gonic/gin"
)

// Get All User Endpoint
func GetFaculties(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from faculty")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_faculty.Faculty{}

	for rows.Next() {
		p := model_faculty.Faculty{}
		err = rows.Scan(&p.ID, &p.Name, &p.ShortName)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetSchools(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from school")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_school.School{}

	for rows.Next() {
		p := model_school.School{}
		err = rows.Scan(&p.ID, &p.Name, &p.FacultyId, &p.ShortName)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetSpecialities(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from specialty")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_specialty.Speciality{}

	for rows.Next() {
		p := model_specialty.Speciality{}
		err = rows.Scan(&p.ID, &p.Name, &p.ShortName, &p.FacultyId, &p.SchoolId)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetCourses(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from curses_number")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_course.Course{}

	for rows.Next() {
		p := model_course.Course{}
		err = rows.Scan(&p.ID, &p.Number)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetGroups(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from groups")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_groupe.Group{}

	for rows.Next() {
		p := model_groupe.Group{}
		err = rows.Scan(&p.ID, &p.Name, &p.FacultyId, &p.SchoolId, &p.SpecialityId)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetLanguages(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from language")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_language.Language{}

	for rows.Next() {
		p := model_language.Language{}
		err = rows.Scan(&p.ID, &p.Name, &p.ShortName)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetLections(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from lection")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_lection.Lection{}

	for rows.Next() {
		p := model_lection.Lection{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}

func GetUserType(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from usertype")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_usertype.UserType{}

	for rows.Next() {
		p := model_usertype.UserType{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"user": &repos,
	})
	defer db.Close()
}
