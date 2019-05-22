package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"time"
)

var (
	LevelTicker = time.Tick(time.Hour)
	LevelTickerSize = 0
)

func baseLevelLogic (screen *ebiten.Image, timelimit time.Duration) {

	if LevelManager.NewState {
		if timelimit != 0 {
			LevelTicker = time.Tick(timelimit)
		} else {
			LevelTicker = nil
		}

		QuestionsAnswered = []bool{}
		updateButtons()
	}

	if answered || tries >= 3 {
		if timelimit != 0 {
			LevelTicker = time.Tick(timelimit)
			LevelTickerSize = 0
		} else {
			LevelTicker = nil
		}

		firstNum, secondNum, operation, answer = generateQuestion(12)
		QuestionsAnswered = append(QuestionsAnswered, answered)
		updateButtons()
		answered = false
		tries = 0
	}

	if len(QuestionsAnswered) > 9 {
		LevelManager.SetLevel("end menu")
		ButtonManager.ClearButtons()
		TextManager.ClearStaticText()
		answered = false
		tries = 0
	}

	if timelimit != 0 {
		ebitenutil.DrawRect(screen, WIDTH/4, HEIGHT/3, WIDTH/2-float64(LevelTickerSize * 9/10), 20, color.RGBA{255, 0, 10, 255})

		select {
		case <-ticker:
			LevelTickerSize++
		case <-LevelTicker:
			firstNum, secondNum, operation, answer = generateQuestion(12)
			QuestionsAnswered = append(QuestionsAnswered, false)
			LevelTickerSize = 0
			updateButtons()
			answered = false
			tries = 0
		default:
		}
	}

	message := fmt.Sprint(firstNum, getSymbol(operation), secondNum)

	RenderQuestionsList(screen, 5, 10)

	TextManager.RenderTextTo("main", message, WIDTH/2, HEIGHT/2, QuestionText, screen)

}