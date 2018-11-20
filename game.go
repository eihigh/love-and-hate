package main

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	scr *ebiten.Image

	game struct {
		state  sio.Stm
		action string
	}
)

type gameScene = int

const (
	title gameScene = iota
	play
	gameOver
)

type gameAction = int

const (
	actionToTitle = iota
	actionNewGame
)

func update(s *ebiten.Image) error {

	scr = s

	if onQuit() {
		return sio.ErrSuccess
	}

	game.action = ""
	game.state.Update()
	switch game.state.Current() {

	case title:
		updateTitle()

	case play:
		updatePlay()

	case gameOver:

	}

	return nil
}

func updateTitle() {

	text := "引き裂くのはたやすかった。\nその勇気はなかった。矛先は自分に向いた"
	tb := sio.NewTextBox(8, text)
	tb.Rect = display.Clone(2, 2).Resize(0, -5*tb.EmHeight)
	bd(tb.Rect)
	drawText(scr, tb, color.White)

	text = "STAGE 1"
	tb = sio.NewTextBox(5, text)
	tb.Rect = display.Clone(5, 5)
	bd(tb.Rect)
	drawText(scr, tb, color.White)

	r := display.Clone(5, 8).Scale(1, 0.25)

	text = "MISSION -- "
	tb = sio.NewTextBox(6, text)
	tb.Rect = r.Clone(4, 4).Scale(0.5, 1)
	bd(tb.Rect)
	drawText(scr, tb, color.White)

	text = "LOVES < 5\nHATES > 50"
	tb = sio.NewTextBox(4, text)
	tb.Rect = r.Clone(6, 6).Scale(0.5, 1)
	bd(tb.Rect)
	drawText(scr, tb, color.White)
}

func updatePlay() {
}

func bd(r *sio.Rect) {
	x, y := r.Pos(7)
	w, h := r.Width(), r.Height()
	ebitenutil.DrawLine(scr, x, y, x+w, y, color.White)
	ebitenutil.DrawLine(scr, x+w, y, x+w, y+h, color.White)
	ebitenutil.DrawLine(scr, x+w, y+h, x, y+h, color.White)
	ebitenutil.DrawLine(scr, x, y+h, x, y, color.White)
}
