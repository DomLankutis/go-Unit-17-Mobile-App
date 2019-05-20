package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

func level0 (screen *ebiten.Image) {

	if LevelManager.NewState {
		QuestionsAnswered = []bool{}
	}

	if answered || LevelManager.NewState {
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

	TextManager.RenderTextTo("main", message, WIDTH/3, HEIGHT/2, QuestionText, screen)

}
