package main

import (
	"./Utils"
	"github.com/hajimehoshi/ebiten"
)

func menu(screen *ebiten.Image) {
	if LevelManager.NewState {
		TextManager.RenderTextTo("title", "Maths Challenge", WIDTH/16, HEIGHT/5, QuestionText, TextManager.StaticTextImage)

		ButtonManager.AddButton("ToLevel0", WIDTH*0.12, HEIGHT*0.35, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
			if state == Utils.Tap {
				LevelManager.SetLevel("level 0")
				ButtonManager.ClearButtons()
				TextManager.ClearStaticText()
			}
		})
		ButtonManager.GetButton("ToLevel0").Img.Fill(ButtonBackground)
		TextManager.RenderTextTo("level selection", "Easy", WIDTH*0.4, HEIGHT*0.4+10, ButtonText, TextManager.StaticTextImage)

		ButtonManager.AddButton("ToLevel1", WIDTH*0.12, HEIGHT*0.5, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
			if state == Utils.Tap {
				LevelManager.SetLevel("level 1")
				ButtonManager.ClearButtons()
				TextManager.ClearStaticText()
			}
		})
		ButtonManager.GetButton("ToLevel1").Img.Fill(ButtonBackground)
		TextManager.RenderTextTo("level selection", "Medium", WIDTH*0.35, HEIGHT*0.55+10, ButtonText, TextManager.StaticTextImage)

		ButtonManager.AddButton("ToLevel2", WIDTH*0.12, HEIGHT*0.65, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
			if state == Utils.Tap {
				LevelManager.SetLevel("level 2")
				ButtonManager.ClearButtons()
				TextManager.ClearStaticText()
			}
		})
		ButtonManager.GetButton("ToLevel2").Img.Fill(ButtonBackground)
		TextManager.RenderTextTo("level selection", "Hard", WIDTH*0.4, HEIGHT*0.7+10, ButtonText, TextManager.StaticTextImage)

	}
}
