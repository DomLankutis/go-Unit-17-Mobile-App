package main

import (
	"github.com/hajimehoshi/ebiten"
	"time"
)

func level1 (screen *ebiten.Image) {
	baseLevelLogic(screen, time.Second * 20)
}
