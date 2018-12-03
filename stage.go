package main

import (
	"image/color"

	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/sio"
)

func c2f(c complex128) (float64, float64) {
	return real(c), imag(c)
}

var (
	red = color.RGBA{255, 0, 0, 255}
)

const (
	phaseTitleIn int = iota
	phaseTitleOut
	phaseBody
	phaseResult
)

type stage struct {
	objs *objects.Objects

	loveIcon, loveBar *sio.Rect
	hateIcon, hateBar *sio.Rect

	states map[string]sio.Stm

	phases     []phase
	phaseIndex int
}

type phase interface {
	update(*stage)
	base() *phaseBase
}

func newStage() *stage {
	s := &stage{}
	s.objs = objects.NewObjects()

	s.loveIcon = &sio.Rect{
		X: 160 - 16,
		Y: 4,
		W: 16,
		H: 16,
	}
	s.hateIcon = &sio.Rect{
		X: 160,
		Y: 4,
		W: 16,
		H: 16,
	}
	s.loveBar = s.loveIcon.Clone(4, 6).SetSize(160-16, 6)
	s.hateBar = s.hateIcon.Clone(6, 4).SetSize(160-16, 6)

	s.phases = append(s.phases, newPhase1())

	s.states = map[string]sio.Stm{
		"stage": sio.Stm{},
		"phase": sio.Stm{},
	}
	return s
}

func (s *stage) update() action {
	o := s.objs
	for _, st := range s.states {
		st.Update()
	}

	ph := s.phases[s.phaseIndex]
	ph.update(s)
	bd(s.loveIcon)
	bd(s.hateIcon)

	o.UpdatePlayer()

	for _, sym := range o.Symbols {
		sym.Update()
	}
	o.Collision(view)

	s.draw()

	return noAction
}

func (s *stage) draw() {
	o := s.objs
	dg := &draw.Group{Dst: scr}

	pl := sprites.Sprites["player"]
	yellow := 1 - 0.8*sio.UWave(s.states["stage"].RatioTo(20))
	pl.Draw(dg, draw.Shift(c2f(o.Player.Pos)), draw.Paint(1, 1, yellow, 1))

	for _, sym := range o.Symbols {
		b := sym.Base()
		spr := sprites.HateSprite
		if b.IsLove {
			spr = sprites.LoveSprite
		}
		spr.Draw(dg, draw.Shift(c2f(b.Pos)), draw.Paint(1, 1, 1, sym.Alpha()))
	}

	// UIs
	pb := s.phases[s.phaseIndex].base()
	show := float64(pb.loves.show)
	ratio := float64(pb.loves.min) / show
	bar := s.loveBar.Clone(6, 6).Scale(ratio, 1)
	dg.DrawRect(bar, red)
	//	bar := s.r.loveBar
	//	show := float64(s.phase.showLoves)
	//
	//	min := float64(s.phase.minLoves) / show
	//	ax, ay := bar.Pos(6)
	//	bx, by := ax-bar.W*min, ay
	//	ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
	//
	//	ratio := float64(o.Player.Loves) / show
	//	ax, ay = bar.Pos(6)
	//	bx, by = ax-bar.W*ratio, ay
	//	ebitenutil.DrawLine(scr, ax, ay, bx, by, color.White)
	//
	//	bar = s.r.hateBar
	//	show = float64(s.phase.showHates)
	//
	//	max := float64(s.phase.maxHates) / show
	//	ax, ay = bar.Pos(4)
	//	bx, by = ax+bar.W*max, ay
	//	ebitenutil.DrawLine(scr, ax, ay, bx, by, color.White)
	//
	//	ratio = float64(o.Player.Hates) / show
	//	ax, ay = bar.Pos(4)
	//	bx, by = ax+bar.W*ratio, ay
	//	ebitenutil.DrawLine(scr, ax, ay, bx, by, red)
	//
	//	alpha := 1.0
	//	if s.phase.minLoves > o.Player.Loves {
	//		alpha = 1 - 0.5*sio.UWave(s.state.RatioTo(40))
	//	}
	//	sprites.LoveSprite.Draw(dg, draw.Shift(s.r.love.Pos(5)), draw.Paint(1, 1, 1, alpha))
	//	sprites.HateSprite.Draw(dg, draw.Shift(s.r.hate.Pos(5)))
	//	bd(s.r.love)
	//	bd(s.r.hate)
}

// ------------------------------------------------------------
//  Symbols
// ------------------------------------------------------------

type up struct {
	objects.SymbolBase
	vec   complex128
	state sio.Stm
}

func newUp() *up {
	u := &up{}
	u.Pos = complex(50, 200)
	u.IsLove = true
	return u
}

func (u *up) Alpha() float64 {
	return u.state.RatioTo(10)
}

func (u *up) Update() {
	u.state.Update()
	u.Pos += complex(0, -1)
}

type up2 struct {
	objects.SymbolBase
	vec   complex128
	state sio.Stm
}

func newUp2() *up2 {
	u := &up2{}
	u.Pos = complex(100, 200)
	u.IsLove = false
	return u
}

func (u *up2) Alpha() float64 {
	return u.state.RatioTo(10)
}

func (u *up2) Update() {
	u.state.Update()
	u.Pos += complex(0, -1)
}
