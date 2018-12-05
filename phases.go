package main

import (
	"math/rand"

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

	nways []*nway
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

	p.nways = []*nway{}
	return p
}

func (p *phase1) update(s *stage) {
	p.w.Count++

	switch p.w.State {
	case "begin":
		if p.w.Count > 180 {
			p.w.Switch("main")
		}

	case "main":
		p.updateMain(s)
	}
}

func (p *phase1) updateMain(s *stage) {
	o := s.objs
	if p.w.Count%60 == 0 {
		p.nways = append(p.nways, &nway{
			pos: complex(360*rand.Float64(), 250),
		})
	}

	for _, n := range p.nways {
		if !n.isDead {
			n.update(o)
		}
	}
}

type nway struct {
	n      int
	pos    complex128
	count  int
	isDead bool
}

func (n *nway) update(o *objects.Objects) {
	n.pos += complex(0, -1)
	n.count++
	if n.count > 120 {
		dir := sio.UnitVector(rand.Float64())
		rot := sio.Rot(6)
		for i := 0; i < 6; i++ {
			o.Symbols = append(o.Symbols, newLinear(n.pos, dir*1.5))
			dir *= rot
		}
		n.isDead = true
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

func newLinear(p, v complex128) *linear {
	l := &linear{
		vec: v,
	}
	l.Pos = p
	return l
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

// ------------------------------------------------------------
//  Effects
// ------------------------------------------------------------

type mark struct {
	objects.EffectBase
}

func newMark() *mark {
	m := &mark{}
	m.Type = objects.EffectRipple
	m.Pos = complex(100, 200)
	return m
}

func (m *mark) Update() {
	m.Pos += complex(0, -1)
}
