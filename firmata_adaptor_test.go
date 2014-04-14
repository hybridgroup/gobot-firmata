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
			me.Board = NewBoard(sp{})
			me.Board.Events = append(me.Board.Events, event{Name: "firmware_query"})
			me.Board.Events = append(me.Board.Events, event{Name: "capability_query"})
			me.Board.Events = append(me.Board.Events, event{Name: "analog_mapping_query"})
		}
		adaptor = new(FirmataAdaptor)
	})

	It("Must be able to Finalize", func() {
		Expect(adaptor.Finalize()).To(Equal(true))
	})
	It("Must be able to Connect", func() {
		Expect(adaptor.Connect()).To(Equal(true))
	})
	It("Must be able to Disconnect", func() {
		Expect(adaptor.Disconnect()).To(Equal(true))
	})
	It("Must be able to Reconnect", func() {
		Expect(adaptor.Reconnect()).To(Equal(true))
	})
})
