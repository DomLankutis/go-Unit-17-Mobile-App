package main

import (
	"github.com/hajimehoshi/ebiten"
)

func endMenu(screen *ebiten.Image) {
	if LevelManager.NewState {
		count := 0
		for _, ans := range QuestionsAnswered {
			if ans {
				count++
			}
		}

		var endText string

		switch {
		case count == 10:
			endText = "Amazing Job!"
		case count > 5:
			endText = "Good Job"
		default:
			endText = "Well Done"
		}

		TextManager.RenderTextTo("title", endText, WIDTH/2, HEIGHT/2, QuestionText, TextManager.StaticTextImage)
	}
}