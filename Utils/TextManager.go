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
	fonts map[string]font.Face
}

func NewTextManager() TextManager{
	return TextManager{
		map[string]font.Face{},
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
