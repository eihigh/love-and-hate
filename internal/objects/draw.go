package objects

import (
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/hajimehoshi/ebiten"
)

func (s *SymbolBase) Draw(scr *ebiten.Image) {
	spr := sprites.HateSprite
	if s.IsLove {
		spr = sprites.LoveSprite
	}

	op := &ebiten.DrawImageOptions{}
	spr.Bring(op, s.Pos)
	scr.DrawImage(spr.Image, op)
	/*
		dg := &draw.Group{screen}
		dg.Draw(spr.Image,
			draw.T(spr.Pos(7)),
			draw.T(c2p(sym.Pos)),
			draw.C(1, 1, 1, alpha),
		)
		spr.Draw(dg,
			draw.Rotate(angle),
			draw.Shift(c2p(sym.Pos)),
			draw.Paint(1, 1, 1, alpha),
		)
	*/
}

type drawGroup struct {
	dst  *ebiten.Image
	c    *ebiten.ColorM
	mode ebiten.CompositeMode
}

func newDrawGroup(dst *ebiten.Image, c *ebiten.ColorM, mode ebiten.CompositeMode) *drawGroup {
	return &drawGroup{
		dst:  dst,
		c:    c,
		mode: mode,
	}
}
