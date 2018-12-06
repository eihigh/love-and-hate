package main

import (
	"math/rand"

	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

type mover struct {
	pos, vec complex128
	count    int
	isDead   bool
}

func (m *mover) update() {
	m.pos += m.vec
	m.count++
}

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

	moversMap map[string][]*mover
}

func newPhase1() *phase1 {
	p := &phase1{}
	p.message = "test\nstring"
	p.love.isPositive = true
	p.love.target = 20
	p.love.shown = 100
	p.hate.isPositive = false
	p.hate.target = 10
	p.hate.shown = 30

	p.w = sio.Worker{
		State: "begin",
	}

	p.moversMap = map[string][]*mover{
		"6ways": {},
	}
	return p
}

func (p *phase1) update(s *stage) {
	p.w.Count++

	switch p.w.State {
	case "begin":
		if p.w.Count > 2 {
			p.w.Switch("main")
		}

	case "main":
		p.updateMain(s)
	}
}

func (p *phase1) updateMain(s *stage) {
	o := s.objs

	movers := p.moversMap["6ways"]
	if p.w.Count%60 == 0 {
		movers = append(movers, &mover{
			pos: complex(360*rand.Float64(), 200),
			vec: complex(0, -1),
		})
	}

	for _, m := range movers {
		m.update()

		if m.count > 100 {
			// spawn!
			dir := sio.UnitVector(rand.Float64())
			rot := sio.Rot(6)
			for i := 0; i < 6; i++ {
				o.Symbols = append(o.Symbols, newLinear(m.pos, dir*1.5))
				dir *= rot
			}
			m.isDead = true
		}
	}

	// cleanup movers
	next := make([]*mover, 0, len(movers))
	for _, m := range movers {
		if !m.isDead {
			next = append(next, m)
		}
	}
	p.moversMap["6ways"] = next
}

func (p *phase1) draw(screen *ebiten.Image) {
	dg := &draw.Group{Dst: screen}
	for _, m := range p.moversMap["6ways"] {
		e := objects.EffectBase{
			Pos:   m.pos,
			Type:  objects.EffectRipple,
			Count: m.count,
		}
		e.Draw(dg)
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
