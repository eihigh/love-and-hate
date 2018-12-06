package main

import (
	"log"

	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var (
	scr *ebiten.Image
)

type game struct {
	worker sio.Worker

	// 以下scene変数は同時に存在する可能性がある
	title *title
	stage *stage
}

func newGame() *game {
	// load resources
	sprites.Load()

	// make instances
	g := &game{
		worker: sio.Worker{
			State: "stage",
		},
		title: nil,
		stage: newStage(),
	}
	return g
}

func (g *game) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if input.OnQuit() {
		return sio.ErrSuccess
	}
	scr = screen

	g.worker.Count++
	switch g.worker.State {
	case "title":
		g.updateTitle()

	case "stage":
		g.updateStage()
	}

	return nil
}

func (g *game) updateTitle() {
	a := g.title.update()

	switch a {
	case gameShowTitle:
		log.Fatal("invalid action: title -> gameShowTitle")
	case gameShowStage:
		g.stage = newStage()
		g.worker.State = "stage"
	}
}

func (g *game) updateStage() {
	a := g.stage.update()

	// 各種遷移処理をここで決定する
	switch a {
	case gameShowTitle:
		g.stage = nil // 破棄
		g.title.reuse()
		g.worker.State = "title"
	}
}
