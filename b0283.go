package b0283

import (
	"math"
	"time"
)

type B0283 struct {
	panMin       int // Minimum pan (left/right) angle in degrees, 0°.
	panMax       int // Maximum pan (left/right) angle in degrees, 180°.
	panPosition  int // Current position in degrees. Fully left is 0°. Fully right is 180°.
	panStep      int // Degrees for each move.
	tiltMax      int // Minimum tilt (up/down) angle in degrees, 180°.
	tiltMin      int // Maximum tilt (up/down) angle in degrees, 180°.
	tiltPosition int // Current position in degrees. Fully down is 0°. Fully up is 180°.
	tiltStep     int // Degrees for each move.
	PanServo     angler
	TiltServo    angler
}

func (b *B0283) PanLeft() (newPos int, err error) {
	nextPos := math.Max(float64(b.panPosition-b.panStep), float64(b.panMin))
	if err := b.PanServo.Angle(int(nextPos)); err != nil {
		return b.panPosition, err
	}
	time.Sleep(10 * time.Millisecond)
	return int(nextPos), nil
}

func (b *B0283) PanRight() (newPos int, err error) {
	nextPos := math.Min(float64(b.panPosition+b.panStep), float64(b.panMax))
	if err := b.PanServo.Angle(int(nextPos)); err != nil {
		return b.panPosition, err
	}
	time.Sleep(10 * time.Millisecond)
	return int(nextPos), nil
}

func (b *B0283) TiltUp() (newPos int, err error) {
	return 0, nil
}

func (b *B0283) TiltDown() (newPos int, err error) {
	return 0, nil
}

type angler interface {
	Angle(int) error
}
