package Utils

import (
	"github.com/hajimehoshi/ebiten"
	"time"
)

const (
	None = iota
	Tap
	Hold
)

type TouchManager struct {
	state int
	tappedOn time.Time
}

func NewTouchManager() TouchManager {
	return TouchManager{
		state: None,
	}
}

func (t *TouchManager) GetTouchPosition(id int) (int, int, int) {
	x, y := ebiten.TouchPosition(id)
	if x != 0 || y != 0 {
		if t.state == None {
			t.state = Tap
			t.tappedOn = time.Now()
		} else if time.Now().After(t.tappedOn.Add(time.Millisecond * 400)) {
			t.state = Hold
		}
	} else {
		t.state = None
	}

	return x, y, t.state
}
