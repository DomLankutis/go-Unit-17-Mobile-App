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
  WIDTH = 720
  HEIGHT = 1280
  SCALE = 1
)

const (
  BSize = 110
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
  ticker = time.Tick(time.Second/10)
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
  answered = false
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
		if b != 0 && float64(a)/float64(b) == float64(a/b) {
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

func GenerateVariationList() []int64 {

	var values []int64

	for ; ; {
		canAdd := true

		n := rand.Int63n(10)

		value := n

		for _, val := range values {
			if value == answer || value == val || value == 0 {
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

	var text string

	correctButton := int(rand.Int63n(6)) + 1
	AnswerVariations := GenerateVariationList()

	for i := 1; i <= 6; i++ {
		x := Spacing*float64(i%3) + WIDTH/5
		y := Spacing*float64(i%2) + HEIGHT - HEIGHT/3

		if i == correctButton {
			text = fmt.Sprint(answer)

			ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
				func(button *Utils.Button, status int) {
					if status == Utils.Tap {
						answered = true
						tries = 0
					}
				})
		} else {
			text = fmt.Sprint(answer + AnswerVariations[i])

			ButtonManager.AddButton(fmt.Sprint(i), x, y, BSize, BSize, "nil",
				func(button *Utils.Button, status int) {
					if status == Utils.Tap {
						button.Img.Fill(color.RGBA{255, 0, 0, 255})
						tries++
					}
				})
		}

		TextManager.RenderTextTo("answer", text, int(x+BSize/2), int(y+BSize/2+BSize/6), ButtonText, TextManager.StaticTextImage)

		ButtonManager.GetButton(fmt.Sprint(i)).Img.Fill(ButtonBackground)
	}
}

func RenderQuestionsList(screen *ebiten.Image, x, y float64) {
	colour := color.RGBA{}

	for i := 0; i < len(QuestionsAnswered); i++ {
		if QuestionsAnswered[i] {
			colour = color.RGBA{0, 255, 0, 255}
		} else {
			colour = color.RGBA{255, 0, 0, 255}
		}
		ebitenutil.DrawRect(screen, float64(i*40)+x, y, 30, 30, colour)
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
	// Change the random generator seed so random numbers differ with every launch of the app
	rand.Seed(time.Now().UnixNano())

	//Adding fonts
	TextManager.AddFont("Roboto-Regular.ttf", "main", truetype.Options{
		Size: 42,
		DPI:  192,
	})

	TextManager.AddFont("Roboto-Regular.ttf", "title", truetype.Options{
		Size: 32,
		DPI:  192,
	})

	TextManager.AddFont("Roboto-Regular.ttf", "answer", truetype.Options{
		Size: 18,
		DPI:  192,
	})

	TextManager.AddFont("Roboto-Regular.ttf", "level selection", truetype.Options{
		Size: 24,
		DPI:  192,
	})

	//Adding music/sounds

	//Adding question level to level Manager
	LevelManager.AddLevel("main menu", menu)
	LevelManager.AddLevel("level 0", level0)
	LevelManager.AddLevel("level 1", level1)
	LevelManager.AddLevel("level 2", level2)
	LevelManager.AddLevel("end menu", endMenu)

	LevelManager.SetLevel("main menu")

	if err := ebiten.Run(update, WIDTH, HEIGHT, SCALE, "Test"); err != nil {
		log.Fatal(err)
	}
}