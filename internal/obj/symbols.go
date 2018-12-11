package obj

import "github.com/eihigh/sio"

type SymbolType int

const (
	_ SymbolType = iota
	SymbolLove
	SymbolHate
)

const (
	BabyTime = 10
)

type SymbolBase struct {
	Type   SymbolType
	Pos    complex128
	IsDead bool
	Timer  sio.Timer
}

func (s *SymbolBase) Base() *SymbolBase { return s }

type Linear struct {
	SymbolBase
	Vec complex128
}

func NewLinear(p, v complex128, t SymbolType) *Linear {
	l := &Linear{
		Vec: v,
	}
	l.Pos = p
	l.Type = t
	return l
}

func (l *Linear) Update() {
	l.Pos += l.Vec
}
