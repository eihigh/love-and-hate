package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/action"
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
)

type play struct {
	isPausing bool
	stage     *stage
}

func newPlay(level int) *play {
	return &play{
		stage: newStage(level),
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
	dg.DrawText("-- PAUSING --", env.View, color.White)

	if input.OnDecide() {
		return action.PlayContinue
	}
	return action.NoAction
}
