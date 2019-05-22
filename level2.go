package main

import (
	"github.com/hajimehoshi/ebiten"
	"time"
)

func level2 (screen *ebiten.Image) {
	baseLevelLogic(screen, time.Second * 10)
}
