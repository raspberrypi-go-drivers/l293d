package l293d_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/l293d"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example() {
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

func ExampleNewL293D() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	fmt.Println(chip)
}

func ExampleL293D_NewMotor() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnable := uint8(12) // PWM compatible
	pin1 := uint8(17)
	pin2 := uint8(27)

	motor1, err := chip.NewMotor(1, pinEnable, pin1, pin2)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	// Looks like motors are not running correctly below some speed percentage
	// I experienced it below 50%
	// Set motor1 speed to 65%
	motor1.SetSpeed(65)
}

func ExampleL293D_GetMotor() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnable := uint8(12) // PWM compatible
	pin1 := uint8(17)
	pin2 := uint8(27)

	_, err = chip.NewMotor(1, pinEnable, pin1, pin2)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	motor1, err := chip.GetMotor(1)
	if err != nil {
		fmt.Println("impossible to get motor1")
	}
	// Looks like motors are not running correctly below some speed percentage
	// I experienced it below 50%
	// Set motor1 speed to 65%
	motor1.SetSpeed(65)
}

func ExampleMotor_SetSpeed() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnable := uint8(12) // PWM compatible
	pin1 := uint8(17)
	pin2 := uint8(27)

	motor1, err := chip.NewMotor(1, pinEnable, pin1, pin2)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	// Looks like motors are not running correctly below some speed percentage
	// I experienced it below 50%
	// Set motor1 speed to 65%
	motor1.SetSpeed(65)
}

func ExampleMotor_EnableBrake() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnable := uint8(12) // PWM compatible
	pin1 := uint8(17)
	pin2 := uint8(27)

	motor1, err := chip.NewMotor(1, pinEnable, pin1, pin2)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	motor1.EnableBrake()
}

func ExampleMotor_DisableBrake() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	chip := l293d.NewL293D()
	pinEnable := uint8(12) // PWM compatible
	pin1 := uint8(17)
	pin2 := uint8(27)

	motor1, err := chip.NewMotor(1, pinEnable, pin1, pin2)
	if err != nil {
		fmt.Println("impossible to create new motor")
	}
	motor1.DisableBrake()
}
