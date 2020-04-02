package model

//Menu
type Menu struct {
	ID      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Child   []int  `json:"child" bson:"child"`
	Url     string `json:"url" bson:"url"`
	Sysname string `json:"sysname" bson:"sysname"`
}

//items of menu
type MenuItems []Menu
