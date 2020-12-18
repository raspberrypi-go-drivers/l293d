# L293D

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/l293d)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/l293d)

The L293D chip allows to control 2 DC motors

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)

## Quick start

```go
import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/l293d"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnableA := uint8(12) // PWM compatible
	pin1a := uint8(17)
	pin2a := uint8(27)

	pinEnableB := uint8(13) // PWM compatible
	pin1b := uint8(23)
	pin2b := uint8(24)
	motor1, err := chip.NewMotor(1, pinEnableA, pin1a, pin2a)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	motor2, err := chip.NewMotor(2, pinEnableB, pin1b, pin2b)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	// Looks like motors are not running correctly below some speed percentage
	// I experienced it below 50%
	// Set motor1 speed to 65%
	motor1.SetSpeed(65)
	// Set motor2 speed to 80%
	motor2.SetSpeed(80)
	time.Sleep(3 * time.Second)
	motor1.SetSpeed(0)
	motor2.SetSpeed(0)
}
```

## Datasheet

L293D datasheet: https://www.ti.com/lit/ds/symlink/l293.pdf

## First glance at L293D

```
           ______________
  1,2EN --|              |-- Vcc1
          |              |
     1A --|              |-- 4A
          |              |
     1Y --|              |-- 4Y
          |              |
    GND --|              |-- GND
          |              |
    GND --|              |-- GND
          |              |
     2Y --|              |-- 3Y
          |              |
     2A --|              |-- 3A
	      |              |
   Vcc2 --|______________|-- 3,4EN
```

  Pins    | Description
  --------|---------------------
  1,2EN   | Enable inputs 1 & 2
  1A      | Input 1 for motor 1
  2A      | Input 2 for motor 1
  1Y & 2Y | To motor 1
  Vcc2    | 5V to 36V input for motors
  3,4EN   | Enable inputs 3 & 4
  3A      | Input 1 for motor 2
  4A      | Input 2 for motor 2
  3Y & 4Y | To motor 2
  Vcc1    | 5V input for chip (can be taken from Raspberry Pi)


  Action (identical for 3,4)   | 1,2EN  | 1A   | 2A    | Comments
  -----------------------------|--------|------|-------|---------------------------------
  Rotation                     | high   | high | low   |
  Reverse rotation             | high   | low  | high  |
  Rotation speed               | PWM    | -    | -     |
  Quick stop                   | low    | -    | -     | 1A and 2A state doesn't matter
  Free rotation                | high   | low  | low   | or high/high


To control speed, a PWM signal have to be sent to pins 1,2EN for motor 1 and 3,4EN for motor 2