package Utils

import (
	"github.com/hajimehoshi/ebiten"
)

type LevelManager struct {
	currentLevel string
	level map[string]func(*ebiten.Image)
}

func NewLevelManager() LevelManager{
	return LevelManager{
		"",
		map[string]func(*ebiten.Image){},
	}
}

func (l *LevelManager) AddLevel(name string, f func(screen *ebiten.Image)) {
	l.level[name] = f
}

func (l *LevelManager) SetLevel(name string) {
	l.currentLevel = name
}

func (l *LevelManager) RunLevel(screen *ebiten.Image) {
	l.level[l.currentLevel](screen)
}