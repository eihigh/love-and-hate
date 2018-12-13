package main

import (
	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
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
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if env.DebugMode && input.OnQuit() {
		return sio.ErrSuccess
	}

	// main logic
	// 各種シーンを状況に合わせて更新・描画する
	g.timers.UpdateAll()
	dg := &draw.Group{}

	pt := g.timers["play"]
	if pt.State != "" {
		g.updatePlay()
	}

	tt := g.timers["title"]
	if tt.State != "" {
		g.updateTitle()

		alpha := 1.0
		then := tt.Do(0, 30, func(t sio.Timer) {
			switch t.State {
			case "in":
				alpha = t.Ratio()
			case "out":
				alpha = 1 - t.Ratio()
			}
		})

		then.Once(func() {
			switch tt.State {
			case "in":
				alpha = 1
				tt.Switch("main")
			case "out":
				alpha = 0
				tt.Switch("")
			}
		})

		dg.DrawImage(g.title.layer, draw.Paint(1, 1, 1, alpha))
	}

	et := g.timers["ending"]
	if et.State != "" {

		alpha := 0.0
		switch et.State {
		case "main":
			alpha = 1

		case "in":
			then := et.Do(0, 120, func(t sio.Timer) {
				alpha = t.Ratio()
			})
			then.Once(func() {
				et.Switch("main")
			})

		case "out":
			then := et.Do(0, 180, func(t sio.Timer) {
				alpha = 1 - t.Ratio()
			})
			then.Once(func() {
				et.Switch("")
			})
		}

		dg.DrawImage(g.ending.layer, draw.Paint(1, 1, 1, alpha))
	}

	return nil
}

func (g *game) updateTitle() {
	a := g.title.update()

	if g.timers["title"].State == "main" {
		switch a {
		case action.StartPlay:
			// g.title は破棄しない
			g.play = newPlay()

			g.timers["title"].Switch("out")
			g.timers["play"].Switch("main")
		}
	}
}

func (g *game) updatePlay() {
	a := g.play.update()

	if g.timers["play"].State == "main" {
		switch a {
		case action.FallbackToTitle:
			g.title = newTitle()

			g.timers["play"].Switch("")
			g.timers["title"].Switch("in")
		}
	}
}
