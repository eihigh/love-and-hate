package main

import (
	"image/color"

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
	timers sio.TimersMap

	// children
	top    *titleTop
	stages *titleStages
}

func newTitle() *title {
	return &title{

		timers: sio.TimersMap{
			"top": &sio.Timer{
				State: "main",
			},
			"stage select": &sio.Timer{},
			"how to play":  &sio.Timer{},
		},

		top:    newTitleTop(),
		stages: newTitleStages(),
	}
}

func (t *title) update() action.Action {

	dg := &draw.Group{}
	t.timers.UpdateAll()

	tt := t.timers["top"]
	st := t.timers["stage select"]

	if tt.State != "" {
		t.top.update()
		x, a := updateLayerTimer(tt)
		dg.DrawImage(t.top.layer, draw.Shift(x, 0), draw.Paint(1, 1, 1, a))

		// scan
		if tt.State == "main" && input.OnDecide() {
			tt.Switch("out")
			st.Switch("in")
		}
	}

	if st.State != "" {
		t.updateStages()
		x, a := updateLayerTimer(st)
		dg.DrawImage(t.stages.layer, draw.Shift(x, 0), draw.Paint(1, 1, 1, a))
	}

	return action.NoAction
}

func updateLayerTimer(t *sio.Timer) (x, alpha float64) {
	dur := 30
	dist := 50.0
	r := t.RatioTo(dur)

	switch t.State {
	case "main":
		return 0, 1

	case "in":
		then := t.Do(0, dur, func(t sio.Timer) {
			x = -dist * (1 - ease.OutQuad(r))
			alpha = r
		})
		then.Once(func() {
			t.Switch("main")
			x, alpha = 0, 1
		})

	case "out":
		then := t.Do(0, dur, func(t sio.Timer) {
			x = dist * ease.InQuad(r)
			alpha = 1 - r
		})
		then.Once(func() {
			t.Switch("")
			x, alpha = dist, 0
		})

	default:
		return dist, 0
	}

	return
}

func (t *title) updateTop() {

	// dg := &draw.Group{}
	//
	// // draw menu
	// for _, m := range t.menus {
	// 	dg.DrawText(m.text, m.rect, obj.White)
	// }
	//
	// // draw cursor
	// ct := t.timers["cursor"]
	// if ct.State == "" {
	// 	x, y := sio.Ctof(t.cursor.a)
	// 	t.cursor.box.Move(x, y)
	// }
	//
	// if ct.State == "moving" {
	// 	then := ct.Do(0, 10, func(tm sio.Timer) {
	// 		pos := t.cursor.a + (t.cursor.b-t.cursor.a)*complex(ease.OutQuad(tm.Ratio()), 0)
	// 		t.cursor.box.Move(sio.Ctof(pos))
	// 	})
	// 	then.Once(func() {
	// 		t.cursor.a = t.cursor.b
	// 		ct.Switch("")
	// 	})
	// }
	//
	// // move cursor
	// if input.OnUp() && t.cursor.index > 0 {
	// 	t.cursor.index--
	// 	ct.Switch("moving")
	// }
	// if input.OnDown() && t.cursor.index < len(t.menus)-1 {
	// 	t.cursor.index++
	// 	ct.Switch("moving")
	// }
	// x, y := t.menus[t.cursor.index].rect.Pos(4)
	// t.cursor.b = complex(x, y)
	//
	// dg.DrawText("--", t.cursor.box, obj.White)
	//
	// // do action
	// if input.OnDecide() {
	// 	return t.menus[t.cursor.index].action
	// }
}

func (t *title) updateStages() {

	tt := t.timers["top"]
	st := t.timers["stage select"]

	if st.State == "main" {
		if input.OnCancel() {
			st.Switch("out")
			tt.Switch("in")
		}
	}

	t.stages.update()
}

// ============================================================
//  タイトルトップ画面
// ============================================================

type titleTop struct {
	layer  *ebiten.Image
	timers sio.TimersMap
	cursor cursor
	menus  []*menuItem
	logo   *sio.Rect
}

func newTitleTop() *titleTop {

	// layout
	mr := env.View.Clone(5, 8).Scale(0.5, 0.4)
	mr0 := mr.Clone(8, 8).SetSize(-1, 20).Drive(5)
	mr1 := mr0.Clone(2, 8).Drive(5)
	x, y := mr0.Pos(4)
	cp := complex(x, y)
	layer, _ := ebiten.NewImage(320, 240, ebiten.FilterDefault)

	return &titleTop{
		layer: layer,

		timers: sio.TimersMap{
			"cursor": &sio.Timer{},
			"logo": &sio.Timer{
				State: "start",
			},
		},

		logo: env.View.Clone(8, 8).Scale(1, 0.6).Drive(5),

		cursor: cursor{
			index: 0,
			box:   sio.NewRect(5, 0, 0, 20, 20),
			a:     cp,
			b:     cp,
		},

		menus: []*menuItem{
			{
				action.ToStages,
				"START",
				mr0,
			},
			{
				action.ToHowTo,
				"HOW TO PLAY",
				mr1,
			},
		},
	}
}

func (t *titleTop) update() {
	t.timers.UpdateAll()
	t.layer.Clear()
	dg := &draw.Group{Dst: t.layer}

	// draw title logo
	lt := t.timers["logo"]
	alpha := 0.0
	if lt.State == "start" {
		then := lt.Do(0, 40, func(t sio.Timer) {
			if t.Count%20 < 7 {
				alpha = 0.4
			}
		})
		then = then.Do(0, 35, func(t sio.Timer) {
			if t.Count%15 < 7 {
				alpha = 0.8
			}
		})
		then.Once(func() {
			lt.State = "main"
		})
	} else {
		alpha = 1
	}

	dg.DrawText("LOVE AND HATE", t.logo, color.Alpha{uint8(255 * alpha)})
}

// ============================================================
//  ステージ選択画面
// ============================================================

type titleStages struct {
	layer  *ebiten.Image
	timers sio.TimersMap
	menus  []*menuItem
	cursor cursor
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

func (t *titleStages) update() {
	t.timers.UpdateAll()
	t.layer.Clear()
	dg := &draw.Group{Dst: t.layer}

	for _, menu := range t.menus {
		dg.DrawText(menu.text, menu.rect, obj.White)
	}
}
