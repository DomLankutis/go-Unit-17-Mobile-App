package main

import (
	"fmt"
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
		Level1Ticker = time.Tick(time.Second * 10)
		QuestionsAnswered = []bool{}
	}

	if answered || LevelManager.NewState{
		firstNum, secondNum, operation, answer = generateQuestion(12)
		QuestionsAnswered = append(QuestionsAnswered, answered)
		updateButtons()
		answered = false
		tries = 0
	} else
	if tries >= 3 {
		firstNum, secondNum, operation, answer = generateQuestion(12)
		QuestionsAnswered = append(QuestionsAnswered, answered)
		updateButtons()
		answered = false
		tries = 0
	}

	if len(QuestionsAnswered) >= 9 {
		LevelManager.SetLevel("")
		ButtonManager.ClearButtons()
		TextManager.ClearStaticText()
		answered = false
		tries = 0
	}

	message := fmt.Sprint(firstNum, getSymbol(operation), secondNum)

	RenderQuestionsList(screen)

	TextManager.RenderTextTo("main", message, WIDTH/4, HEIGHT/2, QuestionText, screen)

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
