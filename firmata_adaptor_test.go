package gobotFirmata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FirmataAdaptor", func() {
	var (
		adaptor *FirmataAdaptor
	)

	BeforeEach(func() {
		connect = func(me *FirmataAdaptor) {
			me.Board = newBoard(sp{})
			me.Board.Events = append(me.Board.Events, event{Name: "firmware_query"})
			me.Board.Events = append(me.Board.Events, event{Name: "capability_query"})
			me.Board.Events = append(me.Board.Events, event{Name: "analog_mapping_query"})
		}
		adaptor = new(FirmataAdaptor)
		adaptor.Connect()
	})

	It("Must be able to Finalize", func() {
		Expect(adaptor.Finalize()).To(Equal(true))
	})
	It("Must be able to Disconnect", func() {
		Expect(adaptor.Disconnect()).To(Equal(true))
	})
	It("Must be able to Reconnect", func() {
		Expect(adaptor.Reconnect()).To(Equal(true))
	})
	It("Must be able to InitServo", func() {
		adaptor.InitServo()
	})
	It("Must be able to ServoWrite", func() {
		adaptor.ServoWrite("1", 50)
	})
	It("Must be able to PwmWrite", func() {
		adaptor.PwmWrite("1", 50)
	})
	It("Must be able to DigitalWrite", func() {
		adaptor.DigitalWrite("1", 1)
	})
	It("DigitalRead should return -1 on no data", func() {
		Expect(adaptor.DigitalRead("1")).To(Equal(-1))
	})
	It("AnalogRead should return -1 on no data", func() {
		Expect(adaptor.AnalogRead("1")).To(Equal(-1))
	})
	It("Must be able to I2cStart", func() {
		adaptor.I2cStart(0x00)
	})
	It("I2cRead should return [] on no data", func() {
		Expect(adaptor.I2cRead(1)).To(Equal(make([]uint16, 0)))
	})
	It("Must be able to I2cWrite", func() {
		adaptor.I2cWrite(make([]uint16, 0))
	})
})
