package main

import (
	"fmt"

	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/love-and-hate/levels/level01"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type level interface {
	Update(screen *ebiten.Image, objs *objects.Objects)
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
	s.level.Update(scr, o)

	s.draw()

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

	o := s.objs

	p := o.Player.Pos
	living := 0

	for _, sym := range o.Symbols {
		alpha := sym.Alpha()
		if alpha < 0.99 {
			continue // skip disabled symbol
		}

		b := sym.Base()
		diff := sio.AbsSq(b.Pos - p)

		if diff < 8*8 {
			if b.IsLove {
				o.Player.Loves++
			} else {
				o.Player.Hates++
			}
			b.IsDead = true
		}

		if !view.Contains(b.Pos) {
			b.IsDead = true
		}

		if !b.IsDead {
			living++
		}
	}

	// clean dead objects
	next := make([]objects.Symbol, 0, living)
	for _, sym := range o.Symbols {
		if !sym.Base().IsDead {
			next = append(next, sym)
		}
	}
}

func (s *stage) draw() {

	o := s.objs
	dg := &draw.Group{Dst: scr}

	spr := sprites.Sprites["player"]
	clr := 0.8 * sio.UWave(s.state.RatioTo(20))
	spr.Draw(dg,
		draw.Shift(c2f(o.Player.Pos)),
		draw.Paint(1, 1, 1-clr, 1),
	)
}

func c2f(c complex128) (float64, float64) {
	return real(c), imag(c)
}
