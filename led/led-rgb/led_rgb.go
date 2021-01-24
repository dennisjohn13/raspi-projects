package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewRgbLedDriver(r, "11", "13", "15")

	work := func() {
		gobot.Every(1*time.Second, func() {
			fmt.Println("Red")
			led.SetRGB(0, 255, 255)
			time.Sleep(3 * time.Second)
			fmt.Println("Green")
			led.SetRGB(255, 0, 255)
			time.Sleep(3 * time.Second)
			fmt.Println("Blue")
			led.SetRGB(255, 255, 0)
			time.Sleep(3 * time.Second)
			fmt.Println("Cyan")
			led.SetRGB(255, 0, 0)
			time.Sleep(3 * time.Second)
			fmt.Println("Purple")
			led.SetRGB(0, 255, 0)
			time.Sleep(3 * time.Second)
			fmt.Println("Yellow")
			led.SetRGB(0, 0, 255)
			time.Sleep(3 * time.Second)
			fmt.Println("White")
			led.SetRGB(0, 0, 0)
			time.Sleep(3 * time.Second)
		})
	}

	robot := gobot.NewRobot("ledBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work)

	robot.Start()
}
