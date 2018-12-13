package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

type title struct {
	timers sio.TimersMap
	layer  *ebiten.Image
	logo   *sio.Rect
}

func newTitle() *title {
	layer, _ := ebiten.NewImage(320, 240, ebiten.FilterDefault)

	return &title{
		logo:  env.View.Clone(8, 8).Scale(1, 0.6).Drive(5),
		layer: layer,

		timers: sio.TimersMap{
			"title": &sio.Timer{},
		},
	}
}

func (t *title) update() action.Action {

	t.layer.Clear()
	dg := &draw.Group{Dst: t.layer}
	t.timers.UpdateAll()

	// draw title logo
	tt := t.timers["title"]
	alpha := 1.0
	then := tt.Do(0, 40, func(t sio.Timer) {
		if t.Count%20 < 7 {
			alpha = 0.4
		} else {
			alpha = 0
		}
	})
	then.Do(0, 35, func(t sio.Timer) {
		if t.Count%15 < 7 {
			alpha = 0.8
		} else {
			alpha = 0
		}
	})

	dg.DrawText("LOVE AND HATE", t.logo, color.Alpha{uint8(255 * alpha)})

	// scan action
	if tt.Count > 30 {
		if input.OnDecide() {
			return action.StartPlay
		}
	}

	return action.NoAction
}
