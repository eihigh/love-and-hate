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
	timer sio.Timer
	layer *ebiten.Image
	logo  *sio.Rect
}

func newTitle() *title {
	return &title{
		logo: env.View.Clone(8, 8).Scale(1, 0.6).Drive(5),
	}
}

func (t *title) update() action.Action {

	t.timer.Update()
	dg := &draw.Group{}

	// draw title logo
	alpha := 1.0

	if t.timer.State == "" {
		then := t.timer.Do(0, 50, func(sio.Timer) {
			alpha = 0
		})
		then = then.Do(0, 50, func(t sio.Timer) {
			if t.Count%20 < 7 {
				alpha = 0.4
			} else {
				alpha = 0
			}
		})
		then = then.Do(0, 35, func(t sio.Timer) {
			if t.Count%15 < 7 {
				alpha = 0.8
			} else {
				alpha = 0
			}
		})
		then.After(20, func(sio.Timer) {
			// scan action
			if input.OnDecide() {
				t.timer.Switch("out")
			}
		})
	}

	if t.timer.State == "out" {
		alpha = 0
		then := t.timer.Do(0, 100, func(t sio.Timer) {
			alpha = 1 - t.Ratio()
		})
		then = then.Do(0, 60, func(sio.Timer) {
			alpha = 0
		})
		then.Once(func() {
			t.timer.Switch("finished")
		})
	}

	if t.timer.State == "finished" {
		return action.StartPlay
	}

	dg.DrawText("LOVE AND HATE", t.logo, color.Alpha{uint8(255 * alpha)})

	return action.NoAction
}
