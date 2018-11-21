package main

import (
	"math/rand"

	"github.com/makki0205/led"
)

func main() {
	led, err := led.NewLED("/dev/tty.usbmodem1421")
	if err != nil {
		panic(err)
	}

	for {
		led.Send(uint8(rand.Int()%255), uint8(rand.Int()%255), uint8(rand.Int()%255))
	}

}
