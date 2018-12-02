package objects

import (
	"github.com/eihigh/love-and-hate/internal/input"
	"github.com/eihigh/sio"
)

type Symbol interface {
	Base() *SymbolBase
	Update()
	Alpha() float64
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
	o := &Objects{}
	o.Player.Pos = complex(160, 120)
	return o
}

func (o *Objects) UpdatePlayer() {
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

	o.Player.Pos += sio.UnitVector(a) * spd
}

func (o *Objects) Collision(view *sio.Rect) {

	p := o.Player.Pos
	living := 0

	for _, sym := range o.Symbols {
		alpha := sym.Alpha()
		if alpha < 0.99 {
			continue // skip disabled symbol
		}

		b := sym.Base()
		diff := sio.AbsSq(b.Pos - p)

		if diff < 8*8 {
			if b.IsLove {
				o.Player.Loves++
			} else {
				o.Player.Hates++
			}
			b.IsDead = true
		}

		if !view.Contains(b.Pos) {
			b.IsDead = true
		}

		if !b.IsDead {
			living++
		}
	}

	// clean dead objects
	next := make([]Symbol, 0, living)
	for _, sym := range o.Symbols {
		if !sym.Base().IsDead {
			next = append(next, sym)
		}
	}

	o.Symbols = next
}
