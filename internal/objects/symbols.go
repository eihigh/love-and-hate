package objects

import "github.com/eihigh/sio"

type SymbolBase struct {
	IsLove bool
	IsDead bool
	Pos    complex128
	State  sio.Stm
}

func (s *SymbolBase) Base() *SymbolBase { return s }

func (s *SymbolBase) UpdateBase(pos complex128) {
	s.Pos = pos
	s.State.Update()
}
