package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"time"
)

var (
	Level1Ticker = time.Tick(time.Second * 20)
	Level1TickerSize = 0
)

func level1 (screen *ebiten.Image) {

	if LevelManager.NewState {
		//Reset ticker
		Level1Ticker = time.Tick(time.Second * 20)
	}

	baseLevelLogic(screen)

	ebitenutil.DrawRect(screen, WIDTH/4, HEIGHT/3, WIDTH/2-float64(Level1TickerSize * 9/10), 20, color.RGBA{255, 0, 10, 255})

	select {
	case <-ticker:
		Level1TickerSize++
	case <-Level1Ticker:
		firstNum, secondNum, operation, answer = generateQuestion(12)
		QuestionsAnswered = append(QuestionsAnswered, false)
		Level1TickerSize = 0
		updateButtons()
		answered = false
		tries = 0
	default:
	}
}
