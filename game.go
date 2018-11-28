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
	stage *stage
}

const (
	sceneOpening int = iota
	sceneTitle
	sceneStage
	sceneResult
)

func (g *game) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	scr = screen

	// update process
	var a action
	switch g.state.Get() {
	case sceneTitle:
		a = g.title.update()
	case sceneStage:
		a = g.stage.update()
	}

	// scan process
	switch a {
	case gameShowTitle:
		g.title.init()
		g.state.To(sceneTitle)
	case gameShowStage:
		g.stage = newStage(g.title.level)
		g.state.To(sceneStage)
	}

	g.state.Update()
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
