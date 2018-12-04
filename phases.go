package main

import (
	"image/color"

	"github.com/eihigh/sio"
)

type emo struct {
	target     int
	shown      int
	isPositive bool
}

func (e *emo) isOver(current int) bool {
	if e.isPositive {
		return false
	}
	return e.target <= current
}

func (e *emo) isPoor(current int) bool {
	if !e.isPositive {
		return false
	}
	return e.target > current
}

func (e *emo) colors() (back, front color.Color) {
	if e.isPositive {
		return red, white
	}
	return white, red
}

func (e *emo) ratios(current int) (back, front float64) {
	s := float64(e.shown)
	back = float64(e.target) / s
	front = float64(current) / s
	return back, front
}

type phaseBase struct {
	message      string
	love, hate   emo
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
