package main

import "github.com/eihigh/sio"

type title struct {
	state  sio.Stm
	cursor int
}

func (t *title) init() {
	t.state.Rebirth()
	// keeps cursor data
}

func (t *title) update() action {
	t.state.Update()

	// scan
	if onUp() {
		return gameShowGameOver
	}

	return noAction
}
