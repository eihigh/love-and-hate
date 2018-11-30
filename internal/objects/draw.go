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
}
