package tables

import (
	"fmt"
	model_student "newsfeeder/httpd/models/student"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateStudent(c *gin.Context) {
	db := PostSQLConfig()

	file, er := c.FormFile("file")
	fmt.Println(file, er)
	user := model_student.Student{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	p := model_student.Student{}
	rows, err := db.Query("select id from students ORDER BY id DESC LIMIT 1")
	for rows.Next() {
		err = rows.Scan(&p.ID)
	}
	index := p.ID + 1
	res, err := db.Exec(" insert into students "+
		"(id,name,surname, fathername,enrollment_year,curs_id,lang_id,group_id, email,lection_id,faculty_id,school_id,specialty_id) "+
		"values ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", index, &user.Name, &user.Surname, &user.Fathername, &user.EnrollmentYear,
		&user.CourseId, &user.LanguageId, &user.GroupId, &user.Email, nil, &user.FacultyId, &user.SchoolId, &user.SpecialityId)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"id":      &index,
			"user":    &res,
		})
	}
	defer db.Close()

}

func CreateMaster(c *gin.Context) {
	db := PostSQLConfig()
	user := model_student.Student{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(user)
	p := model_student.Student{}
	rows, err := db.Query("select id from master ORDER BY id DESC LIMIT 1")
	for rows.Next() {
		err = rows.Scan(&p.ID)
	}
	index := p.ID + 1
	res, err := db.Exec(" insert into master "+
		"(id,name,surname, fathername, email,lection_id,faculty_id,school_id) "+
		"values ($1, $2,$3,$4,$5,$6,$7,$8)", index, &user.Name, &user.Surname, &user.Fathername,
		&user.Email, nil, &user.FacultyId, &user.SchoolId)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"id":      &index,
			"user":    &res,
		})
	}
	defer db.Close()

}
