package Utils

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"log"
)


type TextManager struct {
	fonts           map[string]font.Face
	StaticTextImage *ebiten.Image
}

func NewTextManager(w, h int) TextManager{
	i, _ := ebiten.NewImage(w, h, ebiten.FilterNearest)
	return TextManager{
		map[string]font.Face{},
		i,
	}
}

func (f *TextManager) AddFont(path, name string, options truetype.Options) {
	ff := OpenFile(path)
	defer ff.Close()
	file, err := ioutil.ReadAll(ff)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := truetype.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	f.fonts[name] = truetype.NewFace(tt, &options)
}


func (f *TextManager) RenderTextTo(fontName, message string, x, y int, color color.RGBA, screen *ebiten.Image) {
	text.Draw(screen, message, f.fonts[fontName], x, y, color)
}

func (f *TextManager) ClearStaticText() {
	f.StaticTextImage.Clear()
}

func (f *TextManager) RenderStaticText(screen *ebiten.Image) {
	screen.DrawImage(f.StaticTextImage, &ebiten.DrawImageOptions{})
}