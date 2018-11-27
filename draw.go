package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/objects"
)

type SomeMarker struct {
	pos complex128
	fx  objects.Effect
}

func NewSomeMarker() *SomeMarker {
	return &SomeMarker{
		fx: objects.Effect{
			Type:  objects.EffectRipple,
			Color: color.White,
		},
	}
}

func (s *SomeMarker) Effect() *objects.Effect { return &s.fx }
func (s *SomeMarker) Update() {
	s.pos += 1 + 1i
}

type effect interface {
	Effect() *objects.Effect
	Update()
}

type symbol interface {
	Symbol() *objects.Symbol
	Update()
}
