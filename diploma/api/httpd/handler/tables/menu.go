package tables

import (
	"encoding/json"
	"fmt"
	model_menu "newsfeeder/httpd/models/menu"

	"github.com/gin-gonic/gin"
)

// Get All User Endpoint
func GetAllItems(c *gin.Context) {
	db := PostSQLConfig()
	s := c.Request.URL.Query().Get("user")
	rows, err := db.Query("select * from menu where userinfo = '" + s + "';")
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	repos := model_menu.MenuItems{}
	var childUnit []uint8
	for rows.Next() {
		p := model_menu.Menu{}
		err = rows.Scan(&p.ID, &p.Name, &p.Url, &p.Sysname, &p.UserInfo, &childUnit)
		if err != nil {
			fmt.Println(err)
		}
		s := string(childUnit)

		err := json.Unmarshal([]byte(s), &p.Child)
		if err != nil {
			fmt.Println("error:", err)
		}

		repos = append(repos, p)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"menu": &repos,
	})
	defer db.Close()
}
