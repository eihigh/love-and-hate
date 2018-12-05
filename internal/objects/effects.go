package objects

type EffectType int

const (
	EffectHidden EffectType = iota
	EffectRipple
	EffectCross
)

type EffectBase struct {
	Type   EffectType
	Pos    complex128
	Count  int // basically do not touch from logic
	IsDead bool
}

func (e *EffectBase) Base() *EffectBase { return e }
