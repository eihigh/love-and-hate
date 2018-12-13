package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/obj"
	"github.com/eihigh/sio"
)

var (
	stageTexts = []string{
		"", // dummy
		"CASE 01.\n\nオマツリの灯は遠く",
	}
)

type play struct {
	timers    sio.TimersMap
	isPausing bool
	stage     *stage
	level     int
	message   *sio.Rect
}

func newPlay() *play {
	level := 1
	return &play{
		timers: sio.TimersMap{
			"play": &sio.Timer{},
		},

		stage:   newStage(level),
		level:   level,
		message: env.View.Clone(5, 5).Shift(0, -16),
	}
}

func (p *play) update() action.Action {

	dg := &draw.Group{}
	p.timers.UpdateAll()

	pt := p.timers["play"]

	if pt.State == "" || pt.State == "main" {
		then := pt.Do(50, 150, func(t sio.Timer) {
			dg.DrawText(stageTexts[p.level], p.message, obj.White)
		})
		then.Once(func() {
			pt.Continue("main")
		})
		then.Do(0, 40, func(t sio.Timer) {
			dg.DrawText(stageTexts[p.level], p.message, color.Alpha{uint8(255 * (1 - t.Ratio()))})
		})
	}

	if pt.State == "main" {
		a := p.updateStage()

		if a == action.StageClear {
			p.stage = nil
			// TODO: stop bgm
			p.level++
			if p.level >= len(stageTexts) {
				return action.GameClear
			}

			p.stage = newStage(p.level)
			pt.Switch("")
		}

		if a == action.StageFailed {
			return action.GameOver
		}
	}

	return action.NoAction
}

func (p *play) updateStage() action.Action {

	if p.isPausing {
		a := p.updatePauseMenu()
		if a == action.PlayContinue {
			p.isPausing = false
		}
		return action.NoAction
	}

	// normal playing process
	if input.OnCancel() {
		p.isPausing = true
	}
	return p.stage.update()
}

func (p *play) updatePauseMenu() action.Action {
	dg := &draw.Group{}

	dg.DrawText("** PAUSING **\n\npress ok to continue", p.message, color.White)

	if input.OnDecide() {
		return action.PlayContinue
	}
	return action.NoAction
}
