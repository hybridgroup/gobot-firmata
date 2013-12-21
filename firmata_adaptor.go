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

func (fa *FirmataAdaptor) Connect() bool {
	fa.Board = NewBoard(fa.Port, 57600)
	fa.Board.Connect()
	fa.Connected = true
	return true
}

func (da *FirmataAdaptor) Reconnect() bool  { return false }
func (da *FirmataAdaptor) Disconnect() bool { return false }
func (da *FirmataAdaptor) Finalize() bool   { return false }

func (da *FirmataAdaptor) ServoWrite(pin string, angle byte) {
	p, _ := strconv.Atoi(pin)

	da.Board.SetPinMode(byte(p), SERVO)
	da.Board.AnalogWrite(byte(p), angle)
}

func (da *FirmataAdaptor) PwmWrite(pin string, level byte) {
	p, _ := strconv.Atoi(pin)

	da.Board.SetPinMode(byte(p), PWM)
	da.Board.AnalogWrite(byte(p), level)
}

func (da *FirmataAdaptor) DigitalWrite(pin string, level byte) {
	p, _ := strconv.Atoi(pin)

	da.Board.SetPinMode(byte(p), OUTPUT)
	da.Board.DigitalWrite(byte(p), level)
}

func (da *FirmataAdaptor) DigitalRead(pin string) int {
	p, _ := strconv.Atoi(pin)
	da.Board.SetPinMode(byte(p), INPUT)
	da.Board.TogglePinReporting(byte(p), HIGH, REPORT_DIGITAL)
	da.Board.ReadAndProcess()
	events := da.Board.FindEvents(fmt.Sprintf("digital_read_%v", pin))
	if len(events) > 0 {
		return int(events[len(events)-1].Data[0])
	}
	return -1
}

// NOTE pins are numbered A0-A5, which translate to digital pins 14-19
func (da *FirmataAdaptor) AnalogRead(pin string) int {
	p, _ := strconv.Atoi(pin)
	p = da.digitalPin(p)
	da.Board.SetPinMode(byte(p), ANALOG)
	da.Board.TogglePinReporting(byte(p), HIGH, REPORT_ANALOG)
	da.Board.ReadAndProcess()
	events := da.Board.FindEvents(fmt.Sprintf("analog_read_%v", pin))
	if len(events) > 0 {
		event := events[len(events)-1]
		return int(uint(event.Data[0])<<24 | uint(event.Data[1])<<16 | uint(event.Data[2])<<8 | uint(event.Data[3]))
	}
	return -1
}

func (da *FirmataAdaptor) digitalPin(pin int) int {
	return pin + 14
}

func (fa *FirmataAdaptor) I2cStart(address byte) {
	fa.i2cAddress = address
	fa.Board.I2cConfig([]uint16{0})
}

func (fa *FirmataAdaptor) I2cRead(size uint16) []uint16 {
	fa.Board.I2cReadRequest(fa.i2cAddress, size)
	fa.Board.ReadAndProcess()

	events := fa.Board.FindEvents("i2c_reply")
	if len(events) > 0 {
		return events[len(events)-1].I2cReply["data"]
	}
	return make([]uint16, 0)
}

func (fa *FirmataAdaptor) I2cWrite(data []uint16) {
	fa.Board.I2cWriteRequest(fa.i2cAddress, data)
}
