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
	//Delta x, y between last and current frame
	Dx, Dy int
	ox, oy int
}

func NewTouchManager() TouchManager {
	return TouchManager{
		state: None,
		ox: 0,
		oy: 0,
	}
}

func (t *TouchManager) GetTouchPosition(id int) (int, int, int) {
	x, y := ebiten.TouchPosition(id)
	if x != 0 || y != 0 {
		if t.state == None {
			t.state = Tap
			t.tappedOn = time.Now()
		} else {
			t.state = Hold
		}
	} else {
		t.state = None
	}

	t.Dx = x - t.ox
	t.Dy = y - t.oy

	t.ox, t.oy = x, y

	return x, y, t.state
}
