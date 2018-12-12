package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
	"github.com/fogleman/ease"
	"github.com/hajimehoshi/ebiten"
)

type menuItem struct {
	action action.Action
	text   string
	rect   *sio.Rect
}

type cursor struct {
	a, b  complex128
	box   *sio.Rect
	index int
}

type title struct {
	logo  *sio.Rect
	menus []*menuItem

	cursor cursor

	timers sio.TimersMap

	// children
	top    *titleTop
	stages *titleStages
}

func newTitle() *title {

	// make layouts
	mr := env.View.Clone(5, 8).Scale(0.5, 0.5)
	mr0 := mr.Clone(8, 8).SetSize(-1, 20).Drive(5)
	mr1 := mr0.Clone(2, 8).Drive(5)
	x, y := mr0.Pos(4)
	cp := complex(x, y)

	// make instance
	t := &title{
		timers: sio.TimersMap{
			"title": &sio.Timer{
				State: "top",
			},
			"cursor": &sio.Timer{},
		},

		cursor: cursor{
			index: 0,
			box:   sio.NewRect(5, 0, 0, 20, 20),
			a:     cp,
			b:     cp,
		},

		menus: []*menuItem{
			{
				action.NewGame,
				"NEW GAME",
				mr0,
			},
			{
				action.HowTo,
				"HOW TO PLAY",
				mr1,
			},
		},
	}

	return t
}

func (t *title) update() action.Action {

	t.timers.UpdateAll()
	switch t.timers["title"].State {
	case "top":
		return t.updateTop()
	case "stage select":
	case "how to play":
	}

	return action.NoAction
}

func (t *title) updateTop() action.Action {

	dg := &draw.Group{}

	// draw menu
	for _, m := range t.menus {
		dg.DrawText(m.text, m.rect, obj.White)
	}

	// draw cursor
	ct := t.timers["cursor"]
	if ct.State == "" {
		x, y := sio.Ctof(t.cursor.a)
		t.cursor.box.Move(x, y)
	}

	if ct.State == "moving" {
		then := ct.Do(0, 10, func(tm sio.Timer) {
			pos := t.cursor.a + (t.cursor.b-t.cursor.a)*complex(ease.OutQuad(tm.Ratio()), 0)
			t.cursor.box.Move(sio.Ctof(pos))
		})
		then.Once(func() {
			t.cursor.a = t.cursor.b
			ct.Switch("")
		})
	}

	// move cursor
	if input.OnUp() && t.cursor.index > 0 {
		t.cursor.index--
		ct.Switch("moving")
	}
	if input.OnDown() && t.cursor.index < len(t.menus)-1 {
		t.cursor.index++
		ct.Switch("moving")
	}
	x, y := t.menus[t.cursor.index].rect.Pos(4)
	t.cursor.b = complex(x, y)

	dg.DrawText("--", t.cursor.box, obj.White)

	// do action
	if input.OnDecide() {
		return t.menus[t.cursor.index].action
	}

	return action.NoAction
}

func (t *title) updateStages() action.Action {
	a := t.stages.update()

	switch a {
	case action.Cancel:
		t.timers["stages"].Switch("out")
		t.timers["top"].Switch("in")
		t.top = newTitleTop()

	case action.NewGame:
		t.timers["title"].Switch("leave")
	}

	return action.NoAction
}

// ============================================================
//  タイトルトップ画面
// ============================================================

type titleTop struct{}

func newTitleTop() *titleTop {
	return &titleTop{}
}

// ============================================================
//  ステージ選択画面
// ============================================================

type titleStages struct {
	menus  []*menuItem
	cursor cursor
	layer  *ebiten.Image
	timers sio.TimersMap
}

func newTitleStages() *titleStages {

	// layout
	mr := env.View.Clone(5, 5).Resize(-100, -40)
	mr0 := mr.Clone(8, 8).SetSize(-1, 20).Drive(5)
	mr1 := mr0.Clone(2, 8).Drive(5)
	mr2 := mr1.Clone(2, 8).Drive(5)
	mr3 := mr2.Clone(2, 8).Drive(5)

	layer, _ := ebiten.NewImage(320, 240, ebiten.FilterDefault)

	return &titleStages{
		menus: []*menuItem{
			{
				action.NewGame,
				"CASE 0. aiueo",
				mr0,
			},
			{
				action.NewGame,
				"CASE 1. aiueo",
				mr1,
			},
			{
				action.NewGame,
				"CASE 2. aiueo",
				mr2,
			},
			{
				action.NewGame,
				"CASE 3. aiueo",
				mr3,
			},
		},
		layer: layer,
	}
}

func (t *titleStages) update() action.Action {
	t.timers.UpdateAll()
	return action.NoAction
}
