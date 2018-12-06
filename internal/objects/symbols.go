package objects

type SymbolType int

const (
	SymbolHidden SymbolType = iota
	SymbolLove
	SymbolHate
)

type SymbolBase struct {
	Type   SymbolType
	Pos    complex128
	Count  int
	IsDead bool
}

func (s *SymbolBase) Base() *SymbolBase { return s }
