package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	fmt.Println("Hello World")
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("ledBot", []gobot.Connection{r}, []gobot.Device{led}, work)

	robot.Start()
}
