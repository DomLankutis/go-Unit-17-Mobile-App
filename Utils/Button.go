package Utils

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type Button struct {
	img      *ebiten.Image
	ImgOpt   ebiten.DrawImageOptions
	function func (*Button, int)
	x        float64
	y        float64
	width    float64
	height   float64
}

func (b *Button) GetPosition() (float64, float64) {
	return b.x, b.y
}

func (b *Button) SetPosition(x, y float64) {
	b.x, b.y = x, y
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.ImgOpt.GeoM.SetElement(0, 2, b.x)
	b.ImgOpt.GeoM.SetElement(1, 2, b.y)
	if err := screen.DrawImage(b.img, &b.ImgOpt); err != nil {
		log.Fatal(err)
	}
}

func (b *Button) IsPressed(x, y float64) bool {
	return x >= b.x && b.x + b.width >= x && y >= b.y && b.y + b.height >= y
}

func (b *Button) RunFunc(state int) {
	b.function(b, state)
}
