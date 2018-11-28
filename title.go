package main

import (
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
)

type title struct {
	state  sio.Stm
	cursor int
	level  int
}

func (t *title) init() {
	t.state.Rebirth()
	// keeps cursor data
}

func (t *title) update() action {
	t.state.Update()

	// scan
	if input.OnUp() {
		return gameShowGameOver
	}

	return noAction
}
