package main

import (
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

type ending struct {
	layer  *ebiten.Image
	timers sio.TimersMap
}
