package Utils

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"golang.org/x/mobile/asset"
	"log"
	"strings"
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
	if strings.Contains(path, ".wav") {
		decoded, err := wav.Decode(p.audioContext, file)
		if err != nil {
			log.Fatal(err)
		}
		player, _ := audio.NewPlayer(p.audioContext, decoded)
		p.players[name] = player
	} else if strings.Contains(path, ".mp3") {
		decoded, err := mp3.Decode(p.audioContext, file)
		if err != nil {
			log.Fatal(err)
		}
		player, _ := audio.NewPlayer(p.audioContext, decoded)
		p.players[name] = player
	} else {
		log.Fatal("Can only play mp3 and wav formats")
	}


}

func (p *PlayerManager) Play(name string) {
	p.players[name].Rewind()
	if err := p.players[name].Play(); err != nil {
		log.Fatal(err)
	}
}