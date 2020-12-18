package l293d

import (
	"errors"
	"fmt"
	"math"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	pwmFrequency int = 5000 // 5000Hz
)

// L293D instance
type L293D struct {
	motor1 *Motor
	motor2 *Motor
}

// Motor represents a motor connected to a L293D chip
type Motor struct {
	pinEnable rpio.Pin
	pin1      rpio.Pin
	pin2      rpio.Pin
	brake     bool
}

// NewL293D creates a new L293D instance
func NewL293D() *L293D {
	return &L293D{}
}

// NewMotor creates a new motor instance
func (l293d *L293D) NewMotor(motorID uint, pinEnable uint8, pin1 uint8, pin2 uint8) (*Motor, error) {
	if motorID != 1 && motorID != 2 {
		return nil, errors.New("motorID have to be 1 or 2")
	}
	motor := Motor{
		pinEnable: rpio.Pin(pinEnable),
		pin1:      rpio.Pin(pin1),
		pin2:      rpio.Pin(pin2),
		brake:     false,
	}
	motor.pinEnable.Mode(rpio.Pwm)
	motor.pinEnable.Freq(pwmFrequency)
	motor.pin1.Mode(rpio.Output)
	motor.pin2.Mode(rpio.Output)
	if motorID == 1 {
		l293d.motor1 = &motor
	} else if motorID == 2 {
		l293d.motor2 = &motor
	}
	return &motor, nil
}

// GetMotor returns a pointer to the motor instance motorID
func (l293d *L293D) GetMotor(motorID uint32) (*Motor, error) {
	if motorID != 1 && motorID != 2 {
		return nil, errors.New("motorID have to be 1 or 2")
	}
	if motorID == 1 {
		if l293d.motor1 == nil {
			return nil, fmt.Errorf("motor 1 hasn't been created")
		}
		return l293d.motor1, nil
	}
	if l293d.motor2 == nil {
		return nil, fmt.Errorf("motor 2 hasn't been created")
	}
	return l293d.motor2, nil
}

// SetSpeed set the motor speed as percentage
func (motor *Motor) SetSpeed(speedPct int) {
	if speedPct < -100 {
		speedPct = -100
	} else if speedPct > 100 {
		speedPct = 100
	}
	abs := uint32(math.Abs(float64(speedPct)))
	if speedPct > 0 {
		motor.pinEnable.DutyCycle(abs, 100)
		motor.pin1.High()
		motor.pin2.Low()
	} else if speedPct < 0 {
		motor.pinEnable.DutyCycle(abs, 100)
		motor.pin1.Low()
		motor.pin2.High()
	} else if speedPct == 0 {
		motor.pin1.Low()
		motor.pin2.Low()
		if motor.brake {
			motor.pinEnable.DutyCycle(0, 100)
		} else {
			motor.pinEnable.DutyCycle(100, 100)
		}
	}
}

// EnableBrake enables braking when speed is set to 0%
func (motor *Motor) EnableBrake() {
	motor.brake = true
}

// DisableBrake disables braking when speed is set to 0% (default)
func (motor *Motor) DisableBrake() {
	motor.brake = false
}
