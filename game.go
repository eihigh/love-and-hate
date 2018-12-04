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
	state sio.State

	// 以下scene変数は同時に存在する可能性がある
	title *title
	stage *stage
}

const (
	sceneTitle int = iota
	sceneStage
	sceneResult
)

func newGame() *game {
	sprites.Load()

	g := &game{
		title: nil,
		stage: newStage(),
	}
	g.state.To(sceneStage)
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

	switch g.state.Get() {
	case sceneTitle:
		g.updateTitle()

	case sceneStage:
		g.updateStage()
	}

	g.state.Update()
	return nil
}

func (g *game) updateTitle() {
	a := g.title.update()

	switch a {
	case gameShowTitle:
		log.Fatal("invalid action: title -> gameShowTitle")
	case gameShowStage:
		g.stage = newStage()
		g.state.To(sceneStage)
	}
}

func (g *game) updateStage() {
	a := g.stage.update()

	// 各種遷移処理をここで決定する
	switch a {
	case gameShowTitle:
		g.stage = nil // 破棄
		g.title.reuse()
		g.state.To(sceneTitle)
	}
}
