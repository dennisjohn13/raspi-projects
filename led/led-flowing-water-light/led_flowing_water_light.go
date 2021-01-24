package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

var led [10]*gpio.LedDriver
var device []gobot.Device

func main() {
	pins := [10]string{"29", "31", "7", "10", "11", "12", "13", "15", "16", "18"}
	r := raspi.NewAdaptor()

	for i, pin := range pins {
		led[i] = gpio.NewLedDriver(r, pin)
		device = append(device, led[i])
	}

	//Lights Flowing in a chain
	//work := func() {
	//	gobot.Every(50*time.Millisecond, func() {
	//		for _, l := range led {
	//			time.Sleep(50*time.Millisecond)
	//			l.Toggle()
	//		}
	//	})
	//}

	//Light Flowing Up and Down
	work := func() {
		gobot.Every(50*time.Millisecond, func() {
			for _, l := range led {
				time.Sleep(50 * time.Millisecond)
				l.Toggle()
			}
			for i, _ := range led {
				time.Sleep(50 * time.Millisecond)
				led[len(led)-1-i].Toggle()
			}
		})
	}

	//Single Light Going Up and Down
	//work := func() {
	//	gobot.Every(100*time.Millisecond, func() {
	//		for _, l := range led {
	//			l.On()
	//			time.Sleep(100 * time.Millisecond)
	//			l.Off()
	//		}
	//		for i, _ := range led {
	//			led[len(led)-1-i].On()
	//			time.Sleep(100 * time.Millisecond)
	//			led[len(led)-1-i].Off()
	//
	//		}
	//	})
	//}

	robot := gobot.NewRobot("ledBot",
		[]gobot.Connection{r},
		device,
		work)

	robot.Start()

}
