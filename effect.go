package main

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

type Effect interface {
	Position() complex128
	Info() (EffectType, color.Color, int)
	Update()
}

type BasicEffect struct {
	Pos complex128
}

func (e *BasicEffect) Position() complex128 {
	return e.Pos
}

type RippleEffect struct {
	BasicEffect
	clr   color.Color
	state sio.Stm
}

func (e *RippleEffect) Info() (EffectType, color.Color, int) {
	return EffectRipple, e.clr, e.state.Elapsed()
}

func (e *RippleEffect) Update() {
	e.state.Update()
}
