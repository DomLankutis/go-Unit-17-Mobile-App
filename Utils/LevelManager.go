package Utils

import (
	"github.com/hajimehoshi/ebiten"
)

type LevelManager struct {
	NewState bool
	currentLevel string
	level map[string]func(*ebiten.Image)
}

func NewLevelManager() LevelManager{
	return LevelManager{
		true,
		"",
		map[string]func(*ebiten.Image){},
	}
}

func (l *LevelManager) AddLevel(name string, f func(screen *ebiten.Image)) {
	l.level[name] = f
}

func (l *LevelManager) SetLevel(name string) {
	l.currentLevel = name
	l.NewState = true
}

func (l *LevelManager) RunLevel(screen *ebiten.Image) {
	oldState := l.NewState
	l.level[l.currentLevel](screen)
	if oldState {
		l.NewState = false
	}
}