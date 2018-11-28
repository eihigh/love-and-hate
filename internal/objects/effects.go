package objects

import (
	"image/color"

	"github.com/eihigh/sio"
)

type EffectType int

const (
	_ EffectType = iota
	EffectRipple
	EffectCross
)

type Effect struct {
	Type   EffectType
	Color  color.Color
	Pos    complex128
	State  sio.Stm
	IsDead bool
}

func (e *Effect) Base() *Effect { return e }

func (e *Effect) UpdateBase(pos complex128) {
	e.Pos = pos
	e.State.Update()
}

type RippleState int

// Ripple effect states.
const (
	RippleIn RippleState = iota
	RippleSus
	RippleOut
)