package b0283

import (
	"testing"
)

type ServoMock struct {
}

func (p *ServoMock) Angle(int) error {
	return nil
}

func TestPanLeft(t *testing.T) {
	servo := ServoMock{}

	b := B0283{
		panMax:    180,
		tiltMax:   180,
		panStep:   5,
		PanServo:  &servo,
		TiltServo: &servo,
	}

	want := 0 // Should not move beause initial position is fully left
	got, err := b.PanLeft()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if got != want {
		t.Errorf("Servo position should not move left if already fully left. Got new position %d, want %d", got, want)
	}
}

func TestPanRight(t *testing.T) {
	servo := ServoMock{}
	b := B0283{
		panMax:    180,
		tiltMax:   180,
		panStep:   5,
		PanServo:  &servo,
		TiltServo: &servo,
	}
	want := 5 // Panning right should add [panStep]° to servo position

	got, err := b.PanRight()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if got != want {
		t.Errorf("Got new position %d, want %d", got, want)
	}

	// TODO: Test panning does nothing if next position >= max position
	// TODO: Test Angle() called with the correct args? Or too much in the implementation weeds?
}
