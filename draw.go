package main

import (
	"fmt"
	"image/color"

	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/hajimehoshi/ebiten"
)

var (
	effects = []effect{}
)

type SomeMarker struct {
	objects.Effect
	pos complex128
}

func NewSomeMarker() *SomeMarker {
	s := &SomeMarker{}
	s.Type = objects.EffectRipple
	s.Color = color.NRGBA{0, 0, 255, 255}
	effects = append(effects, s)
	return s
}

func (s *SomeMarker) Update() {
	s.UpdateBase(s.pos)
	s.pos += 1 + 1i
}

type effect interface {
	Base() *objects.Effect
	Update()
}

func drawEffect() {
	for _, effect := range effects {
		b := effect.Base()

		switch b.Type {
		case objects.EffectRipple:
			drawRipple(b)
		}
	}
}

func drawRipple(e *objects.Effect) {
	f := e.State.Age() % 5
	fmt.Printf("f = %+v\n", f)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Apply(e.Color)
	op.GeoM.Translate(real(e.Pos), imag(e.Pos))
}
