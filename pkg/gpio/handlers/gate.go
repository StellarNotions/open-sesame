package handlers

import (
	"time"

	"github.com/StellarNotions/open-sesame/pkg/gpio"
)

func OpenCloseGate(outputPin int) bool {
	pin := gpio.OutputOnPin(outputPin)
	pin.High()
	time.Sleep(time.Second * 2)
	pin.Low()
	return true
}
