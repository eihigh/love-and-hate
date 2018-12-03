package main

import "github.com/eihigh/sio"

type phaseBase struct {
	message      string
	loves, hates struct {
		min, max, show int
	}
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
	p.loves.min = 20
	p.loves.max = 100
	p.loves.show = 100
	p.hates.min = 0
	p.hates.max = 10
	p.hates.show = 20
	return p
}

func (p *phase1) update(s *stage) {
	o := s.objs
	if p.state.HasCounted(7) {
		o.Symbols = append(o.Symbols, newUp())
		p.state.Reset()
	}
	p.state.Update()
}
