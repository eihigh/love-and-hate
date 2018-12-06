package objects

import (
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/sprites"
)

type EffectType int

const (
	EffectHidden EffectType = iota
	EffectRipple
	EffectCross
)

type EffectBase struct {
	Type   EffectType
	Pos    complex128
	Count  int // basically do not touch from logic
	IsDead bool
}

func (e *EffectBase) Base() *EffectBase { return e }

func (e *EffectBase) Draw(dg *draw.Group) {
	switch e.Type {
	case EffectRipple:
		life := 30
		n := e.Count % life
		t := float64(n) / float64(life)
		scale := 0.3 + t
		sprites.Sprites["ripple"].Draw(
			dg,
			draw.Scale(scale, scale),
			draw.Shift(real(e.Pos), imag(e.Pos)),
			draw.Paint(1, 1, 1, 1-t),
		)
	}
}
