package objects

import "github.com/eihigh/sio"

type Symbol struct {
	IsLove bool
	IsDead bool
	Pos    complex128
	State  sio.Stm
}

func (s *Symbol) Update(pos complex128) {
	s.Pos = pos
	s.State.Update()
}
