package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var (
	o *obj.Objects
)

type game struct {
	timers sio.TimersMap

	// scene変数
	title  *title
	play   *play
	ending *ending
}

func newGame() *game {
	// load resources
	images.Load()

	o = &obj.Objects{}
	o.Player.Pos = complex(160, 120)

	return &game{
		timers: sio.TimersMap{
			"title": &sio.Timer{
				State: "main",
			},
			"play":   &sio.Timer{},
			"ending": &sio.Timer{},
		},

		title:  newTitle(),
		play:   nil,
		ending: nil,
	}
}

func (g *game) update() error {
	if env.DebugMode && input.OnQuit() {
		return sio.ErrSuccess
	}

	// main logic
	// 各種シーンを状況に合わせて更新・描画する
	g.timers.UpdateAll()

	pt := g.timers["play"]
	if pt.State == "main" {
		g.updatePlay()
	}

	tt := g.timers["title"]
	if tt.State == "main" {
		g.updateTitle()
	}

	et := g.timers["ending"]
	if et.State == "main" {
		g.updateEnding()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// update player
	dg := &draw.Group{}
	o.UpdatePlayer()

	if o.IsActioned() {
		o.AppendEffect(obj.EffectRippleOnce, o.Player.Pos)
	}

	yellow := 0.8 * sio.UWave(o.Player.Action.RatioTo(20))
	dg.DrawSprite(
		images.Images["player"],
		draw.Shift(sio.Ctof(o.Player.Pos)),
		draw.Paint(1, 1, 1-yellow, 1),
	)

	return nil
}

func (g *game) updateTitle() {
	a := g.title.update()

	switch a {
	case action.StartPlay:
		g.play = newPlay()
		g.timers["title"].Switch("")
		g.timers["play"].Switch("main")
	}
}

func (g *game) updatePlay() {
	a := g.play.update()

	switch a {

	case action.BackToTitle:
		g.title = newTitle()
		g.timers["play"].Switch("")
		g.timers["title"].Switch("main")

	case action.GameClear:
		g.ending = newEnding()
		g.timers["play"].Switch("")
		g.timers["ending"].Switch("main")

	case action.GameOver:
		g.title = newTitle()
		g.timers["play"].Switch("")
		g.timers["title"].Switch("main")

	}
}

func (g *game) updateEnding() {
	a := g.ending.update()

	if a == action.EndingFinished {
		g.title = newTitle()
		g.timers["ending"].Switch("")
		g.timers["title"].Switch("main")
	}
}
