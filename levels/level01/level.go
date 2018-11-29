package level01

import (
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

const (
	starting int = iota
	running
)

type Level struct {
	state sio.Stm
	rot   complex128
}

func New() *Level {
	return &Level{}
}

// Update は各オブジェクトを追加するのみで更新はしない
func (l *Level) Update(scr *ebiten.Image, objs *objects.Objects) {
}
