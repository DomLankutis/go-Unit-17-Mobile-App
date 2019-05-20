package Utils

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type ButtonManager struct {
	buttons map[string]*Button
	ButtonScreen *ebiten.Image
}

func NewButtonManager(width, height int) *ButtonManager {
	buttonScreen, _ :=  ebiten.NewImage(width, height, ebiten.FilterLinear)
	return &ButtonManager{
		map[string]*Button{},
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

func (b *ButtonManager) AddButton(name string, x, y, width, height float64, imgPath string, function func(*Button, int)){
	img, _ := ebiten.NewImage(int(width), int(height), ebiten.FilterLinear)
	if imgPath != "nil" {
		var err error
		img, err = ebiten.NewImageFromImage(OpenImage(imgPath), ebiten.FilterLinear)
		if err != nil {
			log.Fatal(err)
		}
	}


	b.buttons[name] = &Button{
		img,
		ebiten.DrawImageOptions{},
		function,
		x,
		y,
		width,
		height,
	}

	b.updateScreen()
}

func (b *ButtonManager) ClearButtons() {
	b.buttons = map[string]*Button{}
}

func (b *ButtonManager) RemoveButton(name string) {
	if b.buttons[name] != nil {
		b.buttons[name] = nil
	}
}

func (b *ButtonManager) GetButton(name string) *Button {
	if b.buttons[name] != nil {
		return b.buttons[name]
	} else {
		log.Fatal("ButtonManager: No button under name '", name, "' exists")
		return &Button{}
	}
}

func (b *ButtonManager) CheckForPress(x, y, state int) bool{
	var isPressed bool
	for _, button := range b.buttons {
		if state != None && button.IsPressed(float64(x), float64(y)){
			isPressed = true
			button.RunFunc(state)
			break
		}
	}
	b.updateScreen()

	return isPressed
}