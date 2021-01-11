package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "13")
	button := gpio.NewButtonDriver(r, "16")

	//work := func() {
	//	gobot.Every(1*time.Second, func() {
	//		led.Toggle()
	//	})
	//}

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println(gpio.ButtonPush)
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("ledBot",
		[]gobot.Connection{r},
		[]gobot.Device{led, button},
		work)

	robot.Start()
}
