package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/text"
	"github.com/eihigh/sio"
)

type title struct {
	state  sio.Stm
	cursor int
	level  int
}

func newTitle() *title {
	t := &title{}
	return t
}

func (t *title) reuse() {
	t.state.Rebirth()
	// keeps cursor data
}

func (t *title) update() action {

	tb := sio.NewTextBox(5, "LOVE AND HATE\n\nPRESS Z")
	tb.Rect = view.Clone(8, 8).Scale(1, 0.7)
	bd(tb.Rect)
	text.Draw(scr, tb, color.White)

	t.state.Update()

	// scan
	if input.OnUp() {
		return gameShowGameOver
	}

	return noAction
}
