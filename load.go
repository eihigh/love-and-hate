package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	sprites = map[string]*sprite{}

	images = []string{
		"player",
		"love",
		"hate",
		"ripple",
		"cross",
	}
)

func load() {

	// まだローカルから読み込みこむだけ
	for _, img := range images {
		i, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("i/img/%s.png", img), ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		s := newSprite(i)
		sprites[img] = s
	}
}
