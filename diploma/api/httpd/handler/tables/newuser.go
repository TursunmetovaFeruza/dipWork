package tables

import (
	"database/sql"
	"fmt"
	"newsfeeder/config"
	model_user "newsfeeder/httpd/models/user"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Get DB from Mongo Config
func PostSQLConfig() *sql.DB {
	db, err := config.GetPSQL()
	if err != nil {
		fmt.Println("ERROR IS", err)
	}
	return db
}

// Get All User Endpoint
func GetAllUserSQL(c *gin.Context) {
	db := PostSQLConfig()
	rows, err := db.Query("select * from usersInfo")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createTable()
	}
	defer rows.Close()
	repos := []model_user.User{}

	for rows.Next() {
		p := model_user.User{}
		err = rows.Scan(&p.ID, &p.Name, &p.Address, &p.Age, &p.CreatedAt, &p.UpdatedAt)
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

func createTable() {
	db := PostSQLConfig()
	_, err := db.Exec("CREATE TABLE usersInfo (id INTEGER PRIMARY KEY," +
		"name  VARCHAR  NOT NULL," +
		"address VARCHAR  NOT NULL," +
		"age integer NOT NULL," +
		"created_at TIME NOT NULL," +
		"updated_at TIME NOT NULL)")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}

func AddUser(c *gin.Context) {
	db := PostSQLConfig()
	res, err := db.Exec("insert into usersInfo (id,name, address, age,created_at,updated_at) values (2,'iPhone X','iPhone X', $1, $2,$3)",
		32, time.Now(), time.Now())
	if err != nil {
		createTable()
	}
	c.JSON(200, gin.H{
		"message": "Succes Insert User",
		"user":    &res,
	})
	defer db.Close()

}
