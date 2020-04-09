package tables

import (
	"fmt"
	model_user "newsfeeder/httpd/models/signin"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Get All User Endpoint
func GetAllUsers(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := []model_user.SignIn{}

	for rows.Next() {
		p := model_user.SignIn{}
		err = rows.Scan(&p.ID, &p.Username, &p.Password, &p.CreatedAt, &p.UpdatedAt, &p.UserInfo, &p.UserType)
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

func createUserTable() {
	db := PostSQLConfig()
	_, err := db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, username text,password text,created_at TIME NOT NULL, updated_at TIME NOT NULL, user_info_id INTEGER, usertype INTEGER)")
	if err != nil {
		fmt.Println("Problem with creating user table", err)
	}
	defer db.Close()

}

func SignUp(c *gin.Context) {
	db := PostSQLConfig()
	user := model_user.SignIn{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(user)
	stored := model_user.SignIn{}
	check_username := db.QueryRow("select username from users where username=$1", user.Username)
	err = check_username.Scan(&stored.Username)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "Username is exist",
			"user":    &stored.Username,
		})
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	rows, err := db.Query("select id from users ORDER BY id DESC LIMIT 1")
	repos := []model_user.SignIn{}
	for rows.Next() {
		p := model_user.SignIn{}
		err = rows.Scan(&p.ID)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	index := repos[0].ID + 1
	res, err := db.Exec("insert into users (id, username, password,created_at,updated_at,user_info_id,usertype) values ( $1, $2,$3,$4,$5,$6,$7)",
		index, user.Username, hashedPassword, time.Now(), time.Now(), nil, user.UserType)
	if err != nil {
		fmt.Println("create error:", err)
	}
	c.JSON(200, gin.H{
		"message": "Succes",
		"user":    &res,
	})
	defer db.Close()

}

func SignIn(c *gin.Context) {
	db := PostSQLConfig()
	user := model_user.SignIn{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	var usertype string
	stored := model_user.SignIn{}
	res := db.QueryRow("select password, users.id as id, usertype.name as type from users inner join usertype on users.usertype = usertype.id where username=$1", user.Username)
	err = res.Scan(&stored.Password, &stored.UserInfo, &usertype)
	if err != nil {
		fmt.Println("signIn problem:", err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(user.Password)); err != nil {
		c.JSON(200, gin.H{
			"message": "Not correct",
		})
	} else {
		c.JSON(200, gin.H{
			"message":  "Succes",
			"user":     &user.Username,
			"user_id":  &stored.UserInfo,
			"userType": &usertype,
		})
	}
	defer db.Close()

}
