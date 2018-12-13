package obj

import (
	"github.com/eihigh/love-and-hate/internal/env"
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
)

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
		Pos              complex128
		LastLoves, Loves int
		LastHates, Hates int
		Action           sio.Timer
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

func (o *Objects) IsActioned() bool {
	return o.Player.Action.State == "on"
}

func (o *Objects) UpdatePlayer() {

	o.Player.Action.Continue("")
	o.Player.Action.Update()
	if input.JustDecided() {
		o.Player.Action.Continue("on")
	}

	r, l, u, d := input.OnRight(), input.OnLeft(), input.OnUp(), input.OnDown()
	if r && l {
		r, l = false, false
	}
	if u && d {
		u, d = false, false
	}

	// 1直角=1.0
	a := 0.0
	spd := 2.0 + 0i
	if r {
		if u {
			a = -0.5
		} else if d {
			a = 0.5
		} else {
			a = 0.0
		}
	} else if l {
		if u {
			a = -1.5
		} else if d {
			a = 1.5
		} else {
			a = 2.0
		}
	} else if u {
		a = -1.0
	} else if d {
		a = 1.0
	} else {
		spd = 0
	}

	pos := o.Player.Pos + sio.UnitVector(a)*spd
	x, y := real(pos), imag(pos)
	if x < env.View.X {
		x = env.View.X
	}
	if y < env.View.Y {
		y = env.View.Y
	}
	if x > env.View.X+env.View.W {
		x = env.View.X + env.View.W
	}
	if y > env.View.Y+env.View.H {
		y = env.View.Y + env.View.H
	}
	o.Player.Pos = complex(x, y)
}
