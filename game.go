package main

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	scr *ebiten.Image
)

type game struct {
	state sio.Stm
	title title
}

const (
	sceneOpening int = iota
	sceneTitle
	sceneStage
	sceneResult
)

func newGame() *game {
	g := &game{}
	g.state.To(sceneStage)
	return g
}

func (g *game) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	scr = screen

	// update process
	g.state.Update()
	var a action
	switch g.state.Current() {
	case sceneTitle:
		a = g.title.update()
	}

	// scan process
	switch a {
	case gameShowTitle:
		g.title.init()
		g.state.To(sceneTitle)
	}

	return nil
}

func bd(r *sio.Rect) {
	x, y := r.Pos(7)
	w, h := r.Width(), r.Height()
	ebitenutil.DrawLine(scr, x, y, x+w, y, color.White)
	ebitenutil.DrawLine(scr, x+w, y, x+w, y+h, color.White)
	ebitenutil.DrawLine(scr, x+w, y+h, x, y+h, color.White)
	ebitenutil.DrawLine(scr, x, y+h, x, y, color.White)
}
