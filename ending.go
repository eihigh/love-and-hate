package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/audio"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var endingText = `
MUSIC

DOVA SYNDROME




SOUND

効果音ラボ
音人〜On Jin〜





PROGRAM/DESIGN/PLAN

ei1chi






Thank you for Playing.
`

type ending struct {
	layer  *ebiten.Image
	box    *sio.Rect
	scroll float64
	timers sio.TimersMap
}

func newEnding() *ending {
	layer, _ := ebiten.NewImage(320, 240, ebiten.FilterDefault)
	audio.PlayBgm("Scent_of_flowers")

	return &ending{
		timers: sio.TimersMap{
			"ending": &sio.Timer{},
		},
		layer:  layer,
		box:    env.View.Clone(8, 8),
		scroll: 330,
	}
}

func (e *ending) update() action.Action {
	e.scroll -= 0.5
	e.timers.UpdateAll()
	dg := &draw.Group{}
	dg.DrawText(endingText, env.View.Clone(8, 8).Shift(0, e.scroll), obj.White)

	if e.timers["ending"].Count > 3000 {
		return action.EndingFinished
	}
	return action.NoAction
}
