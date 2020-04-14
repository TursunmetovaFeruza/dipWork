package main

import (
	gin "newsfeeder/httpd/routes"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	var s gin.Routes
	s.StartGin()
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(5*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}

/*
	Karena variable s memiliki tipe gin.Server milik server.go
	maka s dapat menjalan kan fungsi StartGin
*/
