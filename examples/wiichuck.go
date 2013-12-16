package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-firmata"
	"github.com/hybridgroup/gobot-i2c"
)

func main() {
	firmata := new(gobotFirmata.FirmataAdaptor)
	firmata.Name = "firmata"
	firmata.Port = "/dev/ttyACM0"

	wiichuck := gobotI2C.NewWiichuck(firmata)
	wiichuck.Name = "wiichuck"

	work := func() {
		go func() {
			for {
				fmt.Println("joystick", gobot.On(wiichuck.Events["joystick"]))
			}
		}()
		go func() {
			for {
				fmt.Println("c", gobot.On(wiichuck.Events["c_button"]))
			}
		}()
		go func() {
			for {
				fmt.Println("z", gobot.On(wiichuck.Events["z_button"]))
			}
		}()
	}

	robot := gobot.Robot{
		Connections: []interface{}{firmata},
		Devices:     []interface{}{wiichuck},
		Work:        work,
	}

	robot.Start()
}
