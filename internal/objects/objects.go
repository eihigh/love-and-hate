package objects

type Symbol interface {
	Base() *SymbolBase
	Update()
}

type Effect interface {
	Base() *EffectBase
	Update()
}

type Objects struct {
	Symbols []Symbol
	Effects []Effect

	Player struct {
		Pos          complex128
		Loves, Hates int
	}
}

func NewObjects() *Objects {
	return &Objects{}
}
