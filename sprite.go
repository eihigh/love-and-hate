package main

import (
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

type sprite struct {
	image *ebiten.Image
	rect  *sio.Rect
}

func newSprite(i *ebiten.Image) *sprite {
	w, h := i.Size()
	return &sprite{
		image: i,
		rect:  sio.NewRect(7, 0, 0, float64(w), float64(h)),
	}
}

func (s *sprite) bring(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.rect.Pos(7))
}
