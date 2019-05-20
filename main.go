package main

import (
  "./Utils"
  "fmt"
  "github.com/golang/freetype/truetype"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
  "image/color"
  "log"
  "math/rand"
  "time"
)

const (
  WIDTH = 360
  HEIGHT = 640
  SCALE = 2
)

const (
  BSize = 50
  Spacing = BSize * 1.5
)

const (
  ADD = iota
  SUB
  DIV
  MUL
)

var (
  QuestionText = color.RGBA{0, 0, 0, 255}
  ButtonText = color.RGBA{0, 0, 0, 255}
  ButtonBackground = color.RGBA{247,211,186, 255}
  Background = color.RGBA{245,239,227, 255}
)

var (
  ButtonManager = Utils.NewButtonManager(WIDTH, HEIGHT)
  TextManager = Utils.NewTextManager(WIDTH, HEIGHT)
  PlayerManager = Utils.NewPlayerManager()
  TouchManager = Utils.NewTouchManager()
  LevelManager = Utils.NewLevelManager()
)

var (
  firstNum int64
  secondNum int64
  operation int64
  answer int64
)

var (
  QuestionsAnswered []bool
)

var (
  tries = 0
  answered = true
)

/*
  Generates 3 random numbers where:
    a - The first number used
    b - The second number used
    o - The arithmetic operation used

  Returns a, b, s, and answer
*/
func generateQuestion(max int64) (a, b, o, ans int64) {
  a = rand.Int63n(max)
  b = rand.Int63n(max)
  o = rand.Int63n(4)

  if a < b {
    a, b = b, a
  }

  switch o {
  case ADD:
    return a, b, o, a + b
  case SUB:
    return a, b, o, a - b
  case DIV:
    if float64(a) / float64(b) == float64(a / b) {
      return a, b, o, a / b
    } else {
      return generateQuestion(max)
    }
  case MUL:
    return a, b, o, a * b
  default:
    return 0, 0, 0, 0
  }
}

func getSymbol(o int64) string {
  switch o {
  case ADD:
    return " + "
  case SUB:
    return " - "
  case DIV:
    return " / "
  case MUL:
    return " * "
  default:
    return " "
  }
}

func GenerateVariatonList() []int64 {

  var values []int64

  for ; ; {
    canAdd := true

    n := rand.Int63n(10)

    value := n

    for _, val := range values {
      if val == value || value == 0 {
        canAdd = false
        break
      }
    }

    if canAdd {
      values = append(values, value)
    }

    if len(values) >= 7 {
      return values
    }
  }

}

func updateButtons() {
  TextManager.ClearStaticText()

  correctButton := int(rand.Int63n(6)) + 1
  AnswerVariations := GenerateVariatonList()

  for i := 1; i <= 6; i++ {
    x := Spacing*float64(i%3) + WIDTH/5
    y := Spacing*float64(i%2) + HEIGHT - HEIGHT/3

    if i == correctButton {

      TextManager.RenderTextTo("answer", fmt.Sprint(answer), int(x+BSize/4), int(y+BSize*0.75),
        ButtonText, TextManager.StaticTextImage)

      ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
        func(button *Utils.Button, status int) {
          if status == Utils.Tap {
            answered = true
            tries = 0
          }
        })
    } else {

      TextManager.RenderTextTo("answer", fmt.Sprint(answer+AnswerVariations[i]), int(x+BSize/4), int(y+BSize*0.75),
        ButtonText, TextManager.StaticTextImage)

      ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
        func(button *Utils.Button, status int) {
          if status == Utils.Tap {
            button.Img.Fill(color.RGBA{255, 0, 0, 255})
            tries++
          }
        })
    }

    ButtonManager.GetButton(fmt.Sprint(i)).Img.Fill(ButtonBackground)
  }
}

func RenderQuestionsList(screen *ebiten.Image) {
  colour := color.RGBA{}

  for i := 0; i < len(QuestionsAnswered); i++ {
    if QuestionsAnswered[i] {
      colour = color.RGBA{0, 255, 0, 255}
    } else {
      colour = color.RGBA{255, 0, 0, 255}
    }
    if i == len(QuestionsAnswered) - 1 {
      colour = color.RGBA{250, 218, 94, 255}
    }
    ebitenutil.DrawRect(screen, float64(i * 15) + 5 , 10, 10, 10, colour)

  }
}

func mainMenu(screen *ebiten.Image) {

  if LevelManager.NewState {
    TextManager.RenderTextTo("title", "Maths Challenge", WIDTH/16, HEIGHT/5, QuestionText, TextManager.StaticTextImage)

    ButtonManager.AddButton("ToLevel0", WIDTH*0.12, HEIGHT*0.35, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
      if state == Utils.Tap {
        LevelManager.SetLevel("Level 0")
        ButtonManager.ClearButtons()
        TextManager.ClearStaticText()
      }
    })
    ButtonManager.GetButton("ToLevel0").Img.Fill(ButtonBackground)
    TextManager.RenderTextTo("level selection", "Easy", WIDTH*0.4, HEIGHT*0.4+10, ButtonText, TextManager.StaticTextImage)

    ButtonManager.AddButton("ToLevel1", WIDTH*0.12, HEIGHT*0.5, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
      if state == Utils.Tap {
        LevelManager.SetLevel("Level 1")
        ButtonManager.ClearButtons()
        TextManager.ClearStaticText()
      }
    })
    ButtonManager.GetButton("ToLevel1").Img.Fill(ButtonBackground)
    TextManager.RenderTextTo("level selection", "Medium", WIDTH*0.35, HEIGHT*0.55+10, ButtonText, TextManager.StaticTextImage)

    ButtonManager.AddButton("ToLevel2", WIDTH*0.12, HEIGHT*0.65, WIDTH*0.75, HEIGHT/10, "nil", func(b *Utils.Button, state int) {
      if state == Utils.Tap {
        LevelManager.SetLevel("Level 2")
        ButtonManager.ClearButtons()
        TextManager.ClearStaticText()
      }
    })
    ButtonManager.GetButton("ToLevel2").Img.Fill(ButtonBackground)
    TextManager.RenderTextTo("level selection", "Hard", WIDTH*0.4, HEIGHT*0.7+10, ButtonText, TextManager.StaticTextImage)

  }
}

func update(screen *ebiten.Image) error {

  screen.Fill(Background)

  if err := screen.DrawImage(ButtonManager.ButtonScreen, &ebiten.DrawImageOptions{}); err != nil {
    log.Fatal(err)
  }

  TextManager.RenderStaticText(screen)


  ButtonManager.CheckForPress(TouchManager.GetTouchPosition(0))
  LevelManager.RunLevel(screen)

  return nil
}


func main() {

  LevelManager.SetLevel("Main Menu")

  // Change the random generator seed so random numbers differ with every launch of the app
  rand.Seed(time.Now().UnixNano())

  //Adding fonts
  TextManager.AddFont("Roboto-Regular.ttf", "main", truetype.Options{
    Size: 42,
    DPI: 72,
  })

  TextManager.AddFont("Roboto-Regular.ttf", "title", truetype.Options{
    Size: 42,
    DPI: 72,
  })

  TextManager.AddFont("Roboto-Regular.ttf", "answer", truetype.Options{
    Size: 14,
    DPI: 72,
  })

  TextManager.AddFont("Roboto-Regular.ttf", "level selection", truetype.Options{
    Size: 32,
    DPI: 72,
  })

  //Adding music
  PlayerManager.NewAudioFromPath("annoyed3.wav", "sick")

  //Adding question level to level Manager
  LevelManager.AddLevel("Main Menu", mainMenu)

  LevelManager.AddLevel("Level 0", func(screen *ebiten.Image) {

    if answered {
      firstNum, secondNum, operation, answer = generateQuestion(10)
      QuestionsAnswered = append(QuestionsAnswered, true)
      updateButtons()
      answered = false
      tries = 0
    }

    if tries >= 3 {
      firstNum, secondNum, operation, answer = generateQuestion(10)
      QuestionsAnswered = append(QuestionsAnswered, false)
      updateButtons()
      answered = false
      tries = 0
    }

    if len(QuestionsAnswered) >= 9 {
      LevelManager.SetLevel("")
      answered = false
      tries = 0
    }

    message := fmt.Sprint(firstNum, getSymbol(operation), secondNum)

    RenderQuestionsList(screen)

    TextManager.RenderTextTo("main", message, WIDTH/3, HEIGHT/2, QuestionText, screen)

  })

  if err := ebiten.Run(update, WIDTH, HEIGHT, SCALE, "Test"); err != nil {
    log.Fatal(err)
  }
}