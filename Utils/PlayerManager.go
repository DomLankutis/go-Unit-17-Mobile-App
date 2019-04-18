package Utils

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"golang.org/x/mobile/asset"
	"log"
)

type PlayerManager struct {
	audioContext *audio.Context
	players      map[string]*audio.Player
}

func NewPlayerManager() PlayerManager {
	audioContext, err := audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
	return PlayerManager {
		audioContext,
		map[string]*audio.Player{},
	}
}

func (p *PlayerManager) NewAudioFromPath(path, name string) {
	file, _ := asset.Open(path)
	decoded, err := wav.Decode(p.audioContext, file)
	if err != nil {
		log.Fatal(err)
	}
	player, _ := audio.NewPlayer(p.audioContext, decoded)
	p.players[name] = player
}

func (p *PlayerManager) Play(name string) {
	p.players[name].Rewind()
	if err := p.players[name].Play(); err != nil {
		log.Fatal(err)
	}
}