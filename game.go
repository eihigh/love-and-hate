package main

import (
	"image/color"
	"log"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	scr *ebiten.Image
)

type game struct {
	state sio.Stm
	title *title
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
		log.Fatal("ohoo")
	}
}

func (g *game) updateStage() {
	a := g.stage.update()

	// 各種遷移処理をここで決定する
	switch a {
	case gameShowTitle:
		g.stage = nil // 破棄
		g.title.init()
		g.state.To(sceneTitle)
	}
}

func bd(r *sio.Rect) {
	x, y := r.Pos(7)
	w, h := r.Width(), r.Height()
	ebitenutil.DrawLine(scr, x, y, x+w, y, color.White)
	ebitenutil.DrawLine(scr, x+w, y, x+w, y+h, color.White)
	ebitenutil.DrawLine(scr, x+w, y+h, x, y+h, color.White)
	ebitenutil.DrawLine(scr, x, y+h, x, y, color.White)
}
