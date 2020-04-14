package routes

import (
	handle_user "newsfeeder/httpd/handler/tables"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()
	// .StaticFile("/favicon.ico", "./resources/favicon.ico")
	api := r.Group("/api")
	{
		api.POST("/UploadFile", handle_user.UploadFile)
		api.GET("/GetAllUsers", handle_user.GetAllUsers)
		api.GET("/GetAllStudents", handle_user.GetAllStudents)
		api.GET("/GetAllMaters", handle_user.GetAllMaters)
		api.POST("/user", handle_user.SignIn)
		api.POST("/CreateStudent", handle_user.CreateStudent)
		api.POST("/CreateMaster", handle_user.CreateMaster)
		api.POST("/users", handle_user.SignUp)
		api.GET("/getMenuItems", handle_user.GetAllItems)
		api.GET("/getUser", handle_user.GetUser)
		api.GET("/GetFaculties", handle_user.GetFaculties)
		api.GET("/GetSchools", handle_user.GetSchools)
		api.GET("/GetSpecialities", handle_user.GetSpecialities)
		api.GET("/GetCourses", handle_user.GetCourses)
		api.GET("/GetGroups", handle_user.GetGroups)
		api.GET("/GetLanguages", handle_user.GetLanguages)
		api.GET("/GetLections", handle_user.GetLections)
		api.GET("/GetUserType", handle_user.GetUserType)

	}
	r.Run("0.0.0.0:5000")
}
