package routes

import (
	handle_user "newsfeeder/httpd/handler/tables"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", handle_user.GetAllUsers)
		api.POST("/user", handle_user.SignIn)
		api.POST("/users", handle_user.SignUp)
		api.GET("/users/:id", handle_user.GetUser)
		api.PUT("/users/:id", handle_user.UpdateUser)
		api.DELETE("/users/:id", handle_user.DeleteUser)
		api.GET("/getMenuItems", handle_user.GetAllItems)
	}
	r.Run("0.0.0.0:5000")
}
