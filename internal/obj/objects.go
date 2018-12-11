package obj

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

func (o *Objects) AppendEffect(t EffectType, p complex128) {
	o.Effects = append(o.Effects, &effectObj{
		EffectBase: EffectBase{
			Type: t,
			Pos:  p,
		},
	})
}
