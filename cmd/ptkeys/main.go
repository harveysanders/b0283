package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
	"github.com/harveysanders/b0283"
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Use arrow keys to pan/tilt platform.")
	fmt.Println("Press ESC or Ctrl+c to quit.")

	// Create new connection to i2c-bus on 1 line with address 0x40.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.New(pca9685.Address, "/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}

	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Sets a single PWM channel 0
	pca0.SetChannel(0, 0, 130)

	servoVertical := pca0.ServoNew(0, nil)   // Servo on channel 0
	servoHorizontal := pca0.ServoNew(1, nil) // Servo on channel 1

	ptPlatform := b0283.B0283{
		PanServo:  servoHorizontal,
		PanStep:   5,
		PanMax:    180,
		TiltServo: servoVertical,
		TiltStep:  5,
		TiltMax:   180,
	}

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			ptPlatform.TiltUp()

		case keyboard.KeyArrowDown:
			ptPlatform.TiltDown()

		case keyboard.KeyArrowLeft:
			ptPlatform.PanLeft()

		case keyboard.KeyArrowRight:
			ptPlatform.PanRight()
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}
	}

}
