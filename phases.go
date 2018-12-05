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
	w     sio.Worker
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

	p.w = sio.Worker{
		State: "begin",
	}
	return p
}

func (p *phase1) update(s *stage) {
	p.w.Count++

	switch p.w.State {
	case "begin":
		// 何もしない
		if p.w.Count > 180 {
			p.w.Switch("main")
		}

	case "main":
		p.updateMain(s)
	}
}

func (p *phase1) updateMain(s *stage) {
	o := s.objs
	if p.w.Count%10 == 0 {
		o.Symbols = append(o.Symbols, newLinear(1+1i))
	}
}

// ------------------------------------------------------------
//  Symbols
// ------------------------------------------------------------

type linear struct {
	objects.SymbolBase
	vec complex128
	age float64
}

func newLinear(v complex128) *linear {
	return &linear{
		vec: v,
	}
}

func (l *linear) Update() {
	l.Pos += l.vec
	l.age++
}

func (l *linear) Alpha() float64 {
	if l.age < 10 {
		return l.age / 10
	}
	return 1
}
