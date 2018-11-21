package led

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestToByte(t *testing.T) {
	fmt.Println(ToByte(255))
}

func TestNewLED(t *testing.T) {

	led, err := NewLED("/dev/tty.usbmodem1421")
	if err != nil {
		panic(err)
	}
	for {
		led.Send(uint8(rand.Int()%255), uint8(rand.Int()%255), uint8(rand.Int()%255))
	}

}
