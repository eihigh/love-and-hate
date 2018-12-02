package objects

type SymbolBase struct {
	IsLove bool
	IsDead bool
	Pos    complex128
}

func (s *SymbolBase) Base() *SymbolBase { return s }
