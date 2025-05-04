package blink

import (
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

func Main() {
	// set LED display
	display := microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	for {
		time.Sleep(time.Millisecond * 1000)
		display.SetPixel(0, 0, microbitmatrix.BrightnessFull)
		time.Sleep(time.Millisecond * 1000)
		display.SetPixel(4, 4, microbitmatrix.BrightnessFull)
		display.Display()
	}
}
