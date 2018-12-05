package main

import (
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/sio"
)

type phaseBase struct {
	message    string
	love, hate emo
}

func (p *phaseBase) base() *phaseBase {
	return p
}

type phase1 struct {
	phaseBase
	state sio.Stm
}

func newPhase1() *phase1 {
	p := &phase1{}
	p.message = "あしたはおまつりだよ！\nねえねえ、ママきいてる？"
	p.love.isPositive = true
	p.love.target = 20
	p.love.shown = 100
	p.hate.isPositive = false
	p.hate.target = 10
	p.hate.shown = 30
	return p
}

func (p *phase1) update(s *stage) {
	o := s.objs
	if p.state.HasCounted(7) {
		o.Symbols = append(o.Symbols, newUp())
		o.Symbols = append(o.Symbols, newUp2())
		p.state.Reset()
	}
	p.state.Update()
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
