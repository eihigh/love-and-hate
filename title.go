package main

import (
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
)

type menuItem struct {
	action action
	rect   *sio.Rect
}

type title struct {
	logo *sio.Rect
	menu []*menuItem

	cursor struct {
		a, b complex128
		item int
	}

	workers sio.Workers
}

func newTitle() *title {

	// make layouts
	mr := view.Clone(5, 8).Scale(0.5, 0.5)
	mr.H = 16
	mr0 := mr.Clone(8, 8).Drive(5)
	mr1 := mr0.Clone(2, 8).Drive(5)

	// make instance
	t := &title{
		workers: sio.Workers{
			"cursor": &sio.Worker{},
		},
		logo: view.Clone(8, 8).Scale(1, 0.7).Drive(5),
		menu: []*menuItem{
			{
				action: titleNewGame,
				rect:   mr0,
			},
			{
				action: titleHowTo,
				rect:   mr1,
			},
		},
	}

	return t
}

func (t *title) reuse() {
}

func (t *title) update() action {

	// 	tb := sio.NewTextBox(5, "LOVE AND HATE\n\nPRESS Z")
	// 	tb.Rect = view.Clone(8, 8).Scale(1, 0.7)
	// 	bd(tb.Rect)
	// 	text.Draw(scr, tb, color.White)

	var pos complex128
	cw := t.workers["cursor"]
	switch cw.State {
	case "moving":
		pos = t.cursor.a + t.cursor.b*ease.OutQuad(cw.T(20))
		if cw.Count > 20 {
			cw.Switch("")
		}
	}

	// scan
	if input.OnDecide() {
		return gameShowStage
	}

	return noAction
}
