// +build arm arm64

package gpio

import (
	"log"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func ledCtrl(num rpio.Pin) {
	log.Println("aaa")
	rpio.Open()
	defer rpio.Close()

	pin := rpio.Pin(num)
	log.Println("ccc = ", pin)

	pin.Output()
	/* FIXME: 1shotだと何故かhighにならない... */
	for ii := 0; ii < 2; ii++ {
		pin.Toggle()
		time.Sleep(200 * time.Millisecond)
	}
}

func LedOn() {
	ledCtrl(4)
}

func LedOff() {
	ledCtrl(17)
}

func LedBlue() {
	ledCtrl(5)
}

func LedGreen() {
	ledCtrl(18)
}

func LedLime() {
	ledCtrl(13)
}

func LedPerple() {
	ledCtrl(12)
}

func LedPink() {
	ledCtrl(22)
}

func LedRed() {
	ledCtrl(27)
}

func LedWhite() {
	ledCtrl(6)
}

func LedYellow() {
	ledCtrl(23)
}
