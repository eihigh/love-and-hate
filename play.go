package main

import (
	"fmt"
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/audio"
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

	bgms = []string{
		"", // dummy
		"Retrospect",
	}
)

type play struct {
	timers    sio.TimersMap
	isPausing bool
	stage     *stage
	message   *sio.Rect
	credits   *sio.Rect

	level  int
	credit int
}

func newPlay() *play {
	level := 1
	return &play{
		timers: sio.TimersMap{
			"play": &sio.Timer{},
		},

		stage:   newStage(level),
		message: env.View.Clone(5, 5).Shift(0, -16),
		credits: env.View.Clone(2, 2).Shift(0, -8),

		level:  level,
		credit: 5,
	}
}

func (p *play) update() action.Action {

	dg := &draw.Group{}
	p.timers.UpdateAll()

	pt := p.timers["play"]

	if pt.State == "" || pt.State == "main" {
		then := pt.Do(50, 150, func(t sio.Timer) {
			dg.DrawText(stageTexts[p.level], p.message, obj.White)
			dg.DrawText(fmt.Sprintf("CREDIT %d", p.credit), p.credits, obj.White)
		})
		then = then.Do(0, 40, func(t sio.Timer) {
			dg.DrawText(stageTexts[p.level], p.message, color.Alpha{uint8(255 * (1 - t.Ratio()))})
			dg.DrawText(fmt.Sprintf("CREDIT %d", p.credit), p.credits, color.Alpha{uint8(255 * (1 - t.Ratio()))})
		})
		then.Once(func() {
			pt.Continue("main")
		})
	}

	if pt.State == "main" {
		a := p.updateStage()

		if a == action.StageClear {
			audio.PauseBgm()
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
			audio.PauseBgm()
			p.credit--
			if p.credit < 0 {
				return action.GameOver
			}
			p.stage = newStage(p.level)
			pt.Switch("")
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
