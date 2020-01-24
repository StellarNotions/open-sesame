package gpio

import (
	"fmt"
	"os"

	KitLog "github.com/go-kit/kit/log"
	"github.com/stianeikeland/go-rpio"
)

func OpenPin(pin int) rpio.Pin {
	logger := KitLog.NewLogfmtLogger(KitLog.NewSyncWriter(os.Stderr))
	logger = KitLog.With(logger, "timestamp", KitLog.DefaultTimestampUTC)
	_ = logger.Log("open_pin", pin)

	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	return rpio.Pin(pin)
}

func OutputOnPin(outputPin int) rpio.Pin {
	logger := KitLog.NewLogfmtLogger(KitLog.NewSyncWriter(os.Stderr))
	logger = KitLog.With(logger, "timestamp", KitLog.DefaultTimestampUTC)
	_ = logger.Log("output_pin", outputPin)

	pin := OpenPin(outputPin)
	pin.Output()
	return pin
}
