package gobotFirmata

import (
	"github.com/choffee/gofirmata"
	"github.com/hybridgroup/gobot"
	"strconv"
)

type FirmataAdaptor struct {
	gobot.Adaptor
	Board *firmata.Board
}

func (fa *FirmataAdaptor) Connect() {
	board, err := firmata.NewBoard(fa.Port, 57600)
	if err != nil {
		panic("Could not setup board")
	}
	fa.Board = board
}

func (da *FirmataAdaptor) DigitalWrite(pin string, level string) {
	p, _ := strconv.Atoi(pin)
	l, _ := strconv.Atoi(level)

	da.Board.SetPinMode(byte(p), firmata.MODE_OUTPUT)
	da.Board.WriteDigital(byte(p), byte(l))
}

func (da *FirmataAdaptor) Disconnect() {
}
