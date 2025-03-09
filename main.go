package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

func main() {
	largeMotor, err := ev3dev.TachoMotorFor("ev3-ports:outB", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left large motor on outB: %v", err)
	}
	err = largeMotor.SetStopAction("brake").Err()
	if err != nil {
		log.Fatalf("failed to set brake stop for left large motor on outB: %v", err)
	}
	largeMotor.SetSpeedSetpoint(70 * largeMotor.MaxSpeed() / 100).Command("run-forever")
	checkErrors(largeMotor)
	time.Sleep(5 * time.Second)
	largeMotor.Command("stop")
}

func checkErrors(devs ...ev3dev.Device) {
	for _, d := range devs {
		err := d.(*ev3dev.TachoMotor).Err()
		if err != nil {
			drv, dErr := ev3dev.DriverFor(d)
			if dErr != nil {
				drv = fmt.Sprintf("(missing driver name: %v)", dErr)
			}
			addr, aErr := ev3dev.AddressOf(d)
			if aErr != nil {
				drv = fmt.Sprintf("(missing port address: %v)", aErr)
			}
			log.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}
