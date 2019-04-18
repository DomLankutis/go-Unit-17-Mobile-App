package main

import (
  "./Utils"
  "fmt"
  "github.com/golang/freetype/truetype"
  "github.com/hajimehoshi/ebiten"
  "image/color"
  "log"
)

const (
  WIDTH = 360
  HEIGHT = 640
  SCALE = 2
)

var (
  Temp = "none"
)

var (
  ButtonManager = Utils.NewButtonManager(WIDTH, HEIGHT)
  PlayerManager = Utils.NewPlayerManager()
  TextManager = Utils.NewTextManager()
  TouchManager = Utils.NewTouchManager()
)

func update(screen *ebiten.Image) error {

  ButtonManager.CheckForPress(TouchManager.GetTouchPosition(0))


  if err := screen.DrawImage(ButtonManager.ButtonScreen, &ebiten.DrawImageOptions{}); err != nil {
    log.Fatal(err)
  }
  TextManager.RenderTextTo("main", Temp, WIDTH / 2, HEIGHT / 2, color.RGBA{255, 255, 255,255}, screen)
  return nil
}

func main() {

  //Adding fonts
  TextManager.AddFont("Roboto-Regular.ttf", "main", truetype.Options{
    Size: 14,
    DPI: 96,
  })

  //Adding music
  PlayerManager.NewAudioFromPath("annoyed3.wav", "sick")

  //Adding buttons
  ButtonManager.AddButton(0, 0, 100, 100, "emoji_sick.png", func(b *Utils.Button, state int){
    Temp = "None"
    if state == Utils.Tap {
      PlayerManager.Play("sick")
      Temp = "Tap"
    } else
    if state == Utils.Hold {
      b.ImgOpt.ColorM.RotateHue(0.1)
      Temp = fmt.Sprintln(float64(TouchManager.Dx), float64(TouchManager.Dy))

      x, y := b.GetPosition()
      b.SetPosition(x + float64(TouchManager.Dx), y + float64(TouchManager.Dy))
    }
  })

  err := ebiten.Run(update, WIDTH, HEIGHT, SCALE, "Test")
  if err != nil {
    log.Fatal(err)
  }
}