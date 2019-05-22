package main

import (
	"./Utils"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math"
	"time"
)

var (
	TotalScore = 0
	counter = 0
	firstHalf = true
	backgroundSound string
	endText string
	medal string
)

func endMenu(screen *ebiten.Image) {
	if LevelManager.NewState {
		LevelTicker = time.Tick(time.Second/2)

		for _, ans := range QuestionsAnswered {
			if ans {
				TotalScore++
			}
		}

		switch {
		case TotalScore == 10:
			endText = "Perfectly Done!"
			backgroundSound = "celebration"
			medal = "GoldMedal"
		case TotalScore > 5:
			endText = "Great Job"
			backgroundSound = "yay"
			medal = "SilverMedal"
		case TotalScore < 3:
			endText = "Better Luck Next time"
			backgroundSound = "crickets"
		default:
			endText = "Well Done"
			backgroundSound = "yay"
			medal = "BronzeMedal"
		}

		ButtonManager.AddButton("StarEmpty", -500, -500, 120, 120, "StarEmpty.png",
			func(b *Utils.Button, s int) {})
		ButtonManager.AddButton("StarFullLeftHalf", -500, -500, 120, 120, "StarFullLeftHalf.png",
			func(b *Utils.Button, s int) {})
		ButtonManager.AddButton("StarFullRightHalf", -500, -500, 120, 120, "StarFullRightHalf.png",
			func(b *Utils.Button, s int) {})
		ButtonManager.AddButton("GoldMedal", -500, -500, 120, 120, "GoldMedal.png",
			func(b *Utils.Button, s int) {})
		ButtonManager.AddButton("SilverMedal", -500, -500, 120, 120, "SilverMedal.png",
			func(b *Utils.Button, s int) {})
		ButtonManager.AddButton("BronzeMedal", -500, -500, 120, 120, "BronzeMedal.png",
			func(b *Utils.Button, s int) {})
		starEmpty := ButtonManager.GetButton("StarEmpty")

		for i := 0; i < 5; i++ {
			starEmpty.SetPosition(144*float64(i)+14, HEIGHT/8)
			starEmpty.Draw(TextManager.StaticTextImage)
		}
	}

	select {
	case<-LevelTicker:
		var fullStarHalf *Utils.Button
		if counter != TotalScore {
			if firstHalf {
				fullStarHalf = ButtonManager.GetButton("StarFullLeftHalf")
				fullStarHalf.SetPosition(144*math.Trunc(float64(counter)/2)+14, HEIGHT/8)
			} else {
				fullStarHalf = ButtonManager.GetButton("StarFullRightHalf")
				fullStarHalf.SetPosition(144*math.Trunc(float64(counter)/2)+74, HEIGHT/8)
			}

			fullStarHalf.Draw(TextManager.StaticTextImage)
			PlayerManager.Play("starDing")

			firstHalf = !firstHalf
			counter++
		} else {
			LevelTicker = nil
			counter = 0
			firstHalf = true
			TextManager.RenderTextTo("level selection", fmt.Sprintln(TotalScore,"/", 10), WIDTH/2, HEIGHT/3, QuestionText, TextManager.StaticTextImage)
			TextManager.RenderTextTo("title", endText, WIDTH/2, HEIGHT/2, QuestionText, TextManager.StaticTextImage)
			TextManager.RenderTextTo("level selection", "Go to Menu", WIDTH/2, HEIGHT*0.9, ButtonText, TextManager.StaticTextImage)

			ButtonManager.AddButton("end->menu", WIDTH/4, HEIGHT/2+HEIGHT/3, WIDTH/2, HEIGHT/10,
				"nil", func(b *Utils.Button, state int) {
					if state == Utils.Tap {
						LevelManager.SetLevel("main menu")
						ButtonManager.ClearButtons()
						TextManager.ClearStaticText()
					}
				})

			ButtonManager.GetButton("end->menu").Img.Fill(ButtonBackground)

			if medal != "" {
				ButtonManager.GetButton(medal).SetPosition(WIDTH/3, HEIGHT*0.55)
				ButtonManager.GetButton(medal).Draw(TextManager.StaticTextImage)
			}

			PlayerManager.Play(backgroundSound)
		}
	default:
	}
}