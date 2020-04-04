package model

//Menu
type Menu struct {
	ID       int       `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	Url      string    `json:"url" bson:"url"`
	Sysname  string    `json:"sysname" bson:"sysname"`
	UserInfo string    `json:"userinfo" bson:"userinfo"`
	Child    []SubMenu `json:"child" bson:"child"`
}

type SubMenu struct {
	Url     string `json:"url" bson:"url"`
	Name    string `json:"name" bson:"name"`
	Sysname string `json:"sysname" bson:"sysname"`
}

//items of menu
type MenuItems []Menu
