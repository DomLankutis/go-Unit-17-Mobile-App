package main

import "github.com/hajimehoshi/ebiten"

func endMenu(screen *ebiten.Image) {
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
		break
	case count > 5:
		endText = "Good Job"
		break
	default:
		endText = "Well Done"
	}

	TextManager.RenderTextTo("title", endText, WIDTH/4, HEIGHT/4, QuestionText, screen)
}