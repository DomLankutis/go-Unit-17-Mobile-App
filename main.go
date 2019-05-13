package main

import (
  "./Utils"
  "fmt"
  "github.com/golang/freetype/truetype"
  "github.com/hajimehoshi/ebiten"
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
func generateQuestion(max int64) (a int64, b int64, o, ans int64) {
  a = rand.Int63n(max)
  b = rand.Int63n(max)
  o = rand.Int63n(4)

  switch o {
  case ADD:
    return a, b, o, a + b
  case SUB:
    if a >= b {
      return a, b, o, a - b
    } else {
      return b, a, o, b - a

    }
  case DIV:
    if b != 0 {
      return a, b, o, a / b
    } else {
      return b, a, o, a / b
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
        color.RGBA{255, 255, 255, 255}, TextManager.StaticTextImage)

      ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
        func(button *Utils.Button, status int) {
          if status == Utils.Tap {
            answered = true
            tries = 0
          }
        })
    } else {

      TextManager.RenderTextTo("answer", fmt.Sprint(answer+AnswerVariations[i]), int(x+BSize/4), int(y+BSize*0.75),
        color.RGBA{255, 255, 255, 255}, TextManager.StaticTextImage)

      ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
        func(button *Utils.Button, status int) {
          if status == Utils.Tap {
            button.ImgOpt.ColorM.Reset()
            button.ImgOpt.ColorM.Translate(255, 0, 0, 255)
            tries++
          }
        })
    }

    ButtonManager.GetButton(fmt.Sprint(i)).ImgOpt.ColorM.Translate(0, 0, 100, 255)
  }
}

func LevelFlowControl() {
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
}

func update(screen *ebiten.Image) error {

  if err := screen.DrawImage(ButtonManager.ButtonScreen, &ebiten.DrawImageOptions{}); err != nil {
    log.Fatal(err)
  }

  TextManager.RenderStaticText(screen)


  ButtonManager.CheckForPress(TouchManager.GetTouchPosition(0))
  LevelManager.RunLevel(screen)

  return nil
}


func main() {

  LevelManager.SetLevel("Level 0")

  // Change the random generator seed so random numbers differ with every launch of the app
  rand.Seed(time.Now().UnixNano())

  //Adding fonts
  TextManager.AddFont("Roboto-Regular.ttf", "main", truetype.Options{
    Size: 24,
    DPI: 144,
  })

  TextManager.AddFont("Roboto-Regular.ttf", "answer", truetype.Options{
    Size: 14,
    DPI: 144,
  })

  //Adding music

  //Adding question level to level Manager
  LevelManager.AddLevel("Level 0", func(screen *ebiten.Image) {

    LevelFlowControl()

    message := fmt.Sprint(firstNum, getSymbol(operation), secondNum)

    TextManager.RenderTextTo("main", message, WIDTH/3, HEIGHT/2, color.RGBA{255, 255, 255, 255}, screen)

  })

  err := ebiten.Run(update, WIDTH, HEIGHT, SCALE, "Test")
  if err != nil {
    log.Fatal(err)
  }
}