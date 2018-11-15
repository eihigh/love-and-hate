package main

import (
	ko "github.com/ei1chi/koromo"
)

type gSimple struct {
	stm *ko.Stm
}

func (g *gSimple) generate(syms []*symbol) {
	s := g.stm
	s.Update()

	if s.Elapsed() == 10 {
		syms = append(syms, newLove())
		s.Reset()
	}
}
