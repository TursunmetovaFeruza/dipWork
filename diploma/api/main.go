package main

import (
	gin "newsfeeder/httpd/routes"
)

func main() {
	var s gin.Routes
	s.StartGin()

}

/*
	Karena variable s memiliki tipe gin.Server milik server.go
	maka s dapat menjalan kan fungsi StartGin
*/
