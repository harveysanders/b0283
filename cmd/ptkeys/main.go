package main

import (
	"log"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
	"github.com/harveysanders/b0283"
)

func main() {
	log.Print("Hello")

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

	servoVert := pca0.ServoNew(0, nil) // Servo on channel 0
	servoLR := pca0.ServoNew(1, nil)   // Servo on channel 1

	ptPlatform := b0283.B0283{
		PanServo:  servoLR,
		TiltServo: servoVert,
	}

	// Sample move
	curPos, err := ptPlatform.PanRight()
	if err != nil {
		log.Fatal("pan error:", err)
	}

	log.Println("Curr pos:", curPos)

}
