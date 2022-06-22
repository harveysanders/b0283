package b0283

import (
	"math"
	"time"
)

type B0283 struct {
	PanMin       int // Minimum pan (left/right) angle in degrees, 0°.
	PanMax       int // Maximum pan (left/right) angle in degrees, 180°.
	panPosition  int // Current position in degrees. Fully left is 0°. Fully right is 180°.
	PanStep      int // Degrees for each move.
	TiltMax      int // Minimum tilt (up/down) angle in degrees, 180°.
	TiltMin      int // Maximum tilt (up/down) angle in degrees, 180°.
	tiltPosition int // Current position in degrees. Fully down is 0°. Fully up is 180°.
	TiltStep     int // Degrees for each move.
	PanServo     angler
	TiltServo    angler
}

func (b *B0283) PanLeft() (newPos int, err error) {
	nextPos := math.Max(float64(b.panPosition-b.PanStep), float64(b.PanMin))
	if err := b.PanServo.Angle(int(nextPos)); err != nil {
		return b.panPosition, err
	}

	time.Sleep(10 * time.Millisecond)
	b.panPosition = int(nextPos)
	return int(nextPos), nil
}

func (b *B0283) PanRight() (newPos int, err error) {
	nextPos := math.Min(float64(b.panPosition+b.PanStep), float64(b.PanMax))
	if err := b.PanServo.Angle(int(nextPos)); err != nil {
		return b.panPosition, err
	}

	b.panPosition = int(nextPos)
	time.Sleep(10 * time.Millisecond)
	return int(nextPos), nil
}

func (b *B0283) TiltUp() (newPos int, err error) {
	nextPos := math.Min(float64(b.tiltPosition+b.TiltStep), float64(b.TiltMax))
	if err := b.TiltServo.Angle(int(nextPos)); err != nil {
		return b.tiltPosition, err
	}

	b.tiltPosition = int(nextPos)
	time.Sleep(10 * time.Millisecond)
	return int(nextPos), nil
}

func (b *B0283) TiltDown() (newPos int, err error) {
	nextPos := math.Max(float64(b.tiltPosition-b.TiltStep), float64(b.TiltMin))
	if err := b.TiltServo.Angle(int(nextPos)); err != nil {
		return b.tiltPosition, err
	}

	b.tiltPosition = int(nextPos)
	time.Sleep(10 * time.Millisecond)
	return int(nextPos), nil
}

type angler interface {
	Angle(int) error
}
