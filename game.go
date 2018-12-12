package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

type game struct {
	scene string

	// scene変数
	title *title
	play  *play
}

func newGame() *game {
	// load resources
	images.Load()

	return &game{
		scene: "title",
		title: newTitle(),
		play:  newPlay(1), // 1 is for debug TODO
	}
}

func (g *game) update() error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if env.DebugMode && input.OnQuit() {
		return sio.ErrSuccess
	}

	switch g.scene {
	case "title":
		g.updateTitle()

	case "play":
		g.updatePlay()

	case "ending":
	}

	return nil
}

func (g *game) updateTitle() {
	a := g.title.update()

	switch a {
	case action.NewGame:
	case action.HowTo:
	}
}

func (g *game) updatePlay() {
	g.play.update()

	// TODO 各種遷移処理をここで決定する
}
