package gobotFirmata

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"strconv"
)

type FirmataAdaptor struct {
	gobot.Adaptor
	Board      *board
	i2cAddress byte
}

func (fa *FirmataAdaptor) Connect() {
	fa.Board = NewBoard(fa.Port, 57600)
	fa.Board.Connect()
}

func (da *FirmataAdaptor) Disconnect() {
}

func (da *FirmataAdaptor) ServoWrite(pin string, angle uint8) {
	p, _ := strconv.Atoi(pin)

	da.Board.SetPinMode(byte(p), SERVO)
	da.Board.AnalogWrite(byte(p), byte(angle))
}

func (da *FirmataAdaptor) PwmWrite(pin string, level uint8) {
	p, _ := strconv.Atoi(pin)

	da.Board.SetPinMode(byte(p), PWM)
	da.Board.AnalogWrite(byte(p), byte(level))
}

func (da *FirmataAdaptor) DigitalWrite(pin string, level string) {
	p, _ := strconv.Atoi(pin)
	l, _ := strconv.Atoi(level)

	da.Board.SetPinMode(byte(p), OUTPUT)
	da.Board.DigitalWrite(byte(p), byte(l))
}

func (da *FirmataAdaptor) DigitalRead(pin string) int {
	p, _ := strconv.Atoi(pin)
	da.Board.SetPinMode(byte(p), INPUT)
	da.Board.TogglePinReporting(byte(p), HIGH, REPORT_DIGITAL)
	da.Board.ReadAndProcess()
	events := da.findEvents(fmt.Sprintf("digital_read_%v", pin))
	if len(events) > 0 {
		return int(events[len(events)-1].Data[0])
	}
	return -1
}

func (fa *FirmataAdaptor) I2cStart(address byte) {
	fa.i2cAddress = address
	fa.Board.I2cConfig([]uint16{0})
}

func (fa *FirmataAdaptor) I2cRead(size uint16) []uint16 {
	fa.Board.I2cReadRequest(fa.i2cAddress, size)
	fa.Board.ReadAndProcess()

	events := fa.findEvents("i2c_reply")
	if len(events) > 0 {
		return events[len(events)-1].I2cReply["data"]
	}
	return make([]uint16, 0)
}

func (fa *FirmataAdaptor) I2cWrite(data []uint16) {
	fa.Board.I2cWriteRequest(fa.i2cAddress, data)
}

func (da *FirmataAdaptor) findEvents(name string) []event {
	ret := make([]event, 0)
	for key, val := range da.Board.Events {
		if val.Name == name {
			ret = append(ret, val)
			da.Board.Events = append(da.Board.Events[:key], da.Board.Events[key+1:]...)
		}
	}
	return ret
}
