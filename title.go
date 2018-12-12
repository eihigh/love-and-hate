package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
)

type menuItem struct {
	action action.Action
	rect   *sio.Rect
}

type title struct {
	logo  *sio.Rect
	menus []*menuItem

	cursor struct {
		a, b  complex128
		index int
	}

	timers sio.TimersMap
}

func newTitle() *title {

	// make layouts
	// mr := view.Clone(5, 8).Scale(0.5, 0.5)
	// mr.H = 16
	// mr0 := mr.Clone(8, 8).Drive(5)
	// mr1 := mr0.Clone(2, 8).Drive(5)

	// make instance
	t := &title{
		timers: sio.TimersMap{
			"cursor": &sio.Timer{},
		},
	}

	return t
}

func (t *title) reuse() {
}

func (t *title) update() action.Action {

	t.timers.UpdateAll()

	last := t.cursor.index
	if input.OnDown() && t.cursor.index > 0 {
		t.cursor.index--
		t.cursor.b = complex(t.menus[t.cursor.index].rect.Pos(4))
	}

	ct := t.timers["cursor"]
	if ct.State == "moving" {
		then := ct.Do(0, 10, func(t sio.Timer) {
			pos := t.cursor.a + (t.cursor.b-t.cursor.a)*ease.OutQuad(t.Ratio())
		})
		then.Once(func() {
			t.cursor.a = t.cursor.b
			ct.Switch("")
		})
	}

	return noAction
}
