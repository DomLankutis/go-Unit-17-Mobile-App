package Utils

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type ButtonManager struct {
	buttons []*Button
	ButtonScreen *ebiten.Image
}

func NewButtonManager(width, height int) *ButtonManager {
	buttonScreen, _ :=  ebiten.NewImage(width, height, ebiten.FilterLinear)
	return &ButtonManager{
		nil,
		buttonScreen,
	}
}

func (b *ButtonManager) updateScreen() {
	if err := b.ButtonScreen.Clear(); err != nil {
		log.Fatal(err)
	}

	for _, button := range b.buttons {
		button.Draw(b.ButtonScreen)
	}
}

func (b *ButtonManager) AddButton(x, y, width, height float64, imgPath string, function func(*Button, int)){
	img, err := ebiten.NewImageFromImage(OpenImage(imgPath), ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}

	b.buttons = append(b.buttons, &Button{
		img,
		ebiten.DrawImageOptions{},
		function,
		x,
		y,
		width,
		height,
	})

	b.updateScreen()
}

func (b *ButtonManager) CheckForPress(x, y, state int) {
	for _, button := range b.buttons {
		if state != None && button.IsPressed(float64(x), float64(y)){
			button.RunFunc(state)
			break
		}
	}
	b.updateScreen()
}