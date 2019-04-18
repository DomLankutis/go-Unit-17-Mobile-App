package Utils

import (
	"github.com/hajimehoshi/ebiten"
)

type Button struct {
	img      *ebiten.Image
	ImgOpt   ebiten.DrawImageOptions
	function func (*Button, int)
	x        int
	y        int
	width    int
	height   int
	Changed bool
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.Changed = false

	screen.DrawImage(b.img, &b.ImgOpt)
}

func (b *Button) IsPressed(x, y int) bool {
	return x >= b.x && b.x + b.width >= x && y >= b.y && b.y + b.height >= y
}

func (b *Button) RunFunc(state int) {
	b.function(b, state)
}
