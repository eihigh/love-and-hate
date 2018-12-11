package obj

import (
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/images"
	"github.com/eihigh/sio"
)

type EffectType int

const (
	EffectHidden EffectType = iota
	EffectRipple
	EffectRippleOnce
)

type EffectBase struct {
	Type   EffectType
	Pos    complex128
	IsDead bool
	Timer  sio.Timer
}

func (e *EffectBase) Base() *EffectBase { return e }

func (e *EffectBase) Draw(dg *draw.Group) {
	switch e.Type {
	case EffectRipple, EffectRippleOnce:
		e.drawRipple(dg)
	}
}

func (e *EffectBase) drawRipple(dg *draw.Group) {

	t := e.Timer
	then := t.Do(0, 30, func(t sio.Timer) {
		r := t.Ratio()
		scale := 0.3 + r

		dg.DrawSprite(
			images.Images["ripple"],
			draw.Scale(scale, scale),
			draw.Shift(sio.Ctof(e.Pos)),
			draw.Paint(1, 1, 1, 1-r),
		)
	})

	then.Once(func() {
		if e.Type == EffectRippleOnce {
			e.IsDead = true
		} else {
			t.Count = 0
		}
	})
}

type effectObj struct {
	EffectBase
}

func (e *effectObj) Update() {}
