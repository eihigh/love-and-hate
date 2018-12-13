package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
)

type play struct {
	isPausing bool
	stage     *stage
	level     int
	message   *sio.Rect
}

func newPlay() *play {
	level := 1
	return &play{
		stage:   newStage(level),
		level:   level,
		message: env.View.Clone(5, 5).Shift(0, -16),
	}
}

func (p *play) update() action.Action {

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

	dg.DrawText("** PAUSING **\n\nPress ok to continue", p.message, color.White)

	if input.OnDecide() {
		return action.PlayContinue
	}
	return action.NoAction
}
