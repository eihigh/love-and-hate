package main

import (
	"fmt"

	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/love-and-hate/levels/level01"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type level interface {
	Update(objs *objects.Objects)
	Draw(scr *ebiten.Image)
}

var (
	levelMakers = map[int]func() level{
		1: func() level { return level01.New() },
	}
)

// stageはステージ突入ごとに必ず作り直し
type stage struct {
	state sio.Stm
	objs  *objects.Objects
	level level

	tl sio.Timeline
}

func newStage(level int) *stage {
	tl := sio.Timeline{}
	tl.Append("stage title", 50)
	tl.Append("player fadein", 20)
	tl.Append("play", -1)

	return &stage{
		objs:  objects.NewObjects(),
		level: levelMakers[level](),
		tl:    tl,
	}
}

func (s *stage) update() action {
	o := s.objs
	if debugMode {
		pl := o.Player
		dmsg := fmt.Sprintf("L %d, H %d, len %d, FPS %d", pl.Loves, pl.Hates, len(o.Symbols), int(ebiten.CurrentFPS()))
		ebitenutil.DebugPrint(scr, dmsg)
	}

	switch s.tl.Current() {
	case "stage title":
	case "player fadein":
	case "play":
	}

	s.tl.Update()

	s.movePlayer()

	// call level-specific process
	s.level.Update(o)

	// update symbols
	for _, sym := range o.Symbols {
		sym.Update()
	}

	// collision check
	s.collision()

	s.level.Draw(scr)

	// draw UIs
	s.drawUI()

	s.state.Update()
	return noAction
}

func (s *stage) movePlayer() {
	r, l, u, d := input.OnRight(), input.OnLeft(), input.OnUp(), input.OnDown()
	if r && l {
		r, l = false, false
	}
	if u && d {
		u, d = false, false
	}

	// 1直角=1.0
	a := 0.0
	spd := 2.0 + 0i
	if r {
		if u {
			a = -0.5
		} else if d {
			a = 0.5
		} else {
			a = 0.0
		}
	} else if l {
		if u {
			a = -1.5
		} else if d {
			a = 1.5
		} else {
			a = 2.0
		}
	} else if u {
		a = -1.0
	} else if d {
		a = 1.0
	} else {
		spd = 0
	}

	s.objs.Player.Pos += sio.UnitVector(a) * spd
}

func (s *stage) collision() {
}

func (s *stage) drawUI() {

	o := s.objs
	op := &ebiten.DrawImageOptions{}

	reop(op)
	spr := sprites.Sprites["player"]
	spr.Bring(op, o.Player.Pos)

	clr := 0.8 * sio.UWave(s.state.RatioTo(20))
	op.ColorM.Scale(1, 1, 1-clr, 1)
	scr.DrawImage(spr.Image, op)
}

func reop(op *ebiten.DrawImageOptions) {
	op.GeoM.Reset()
	op.ColorM.Reset()
}
