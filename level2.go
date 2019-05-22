package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"time"
)

var (
	Level2Ticker = time.Tick(time.Second * 10)
	Level2TickerSize = 0
)

func level2 (screen *ebiten.Image) {

	if LevelManager.NewState {
		//Reset Ticker
		Level2Ticker = time.Tick(time.Second * 10)
	}

	baseLevelLogic(screen)

	ebitenutil.DrawRect(screen, WIDTH/4, HEIGHT/3, WIDTH/2-float64(Level2TickerSize * 18 / 10), 20, color.RGBA{255, 0, 10, 255})

	select {
	case <-ticker:
		Level2TickerSize++
	case <-Level2Ticker:
		firstNum, secondNum, operation, answer = generateQuestion(12)
		QuestionsAnswered = append(QuestionsAnswered, false)
		Level2TickerSize = 0
		updateButtons()
		answered = false
		tries = 0
	default:
	}
}
