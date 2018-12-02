package level01

import (
	"github.com/eihigh/love-and-hate/internal/draw"
	"github.com/eihigh/love-and-hate/internal/objects"
	"github.com/eihigh/love-and-hate/internal/sprites"
	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

const (
	starting int = iota
	running
)

type Level struct {
	state sio.Stm
	rot   complex128
}

func New() *Level {
	return &Level{}
}

// Update は各オブジェクトを追加するのみで更新はしない
func (l *Level) Update(screen *ebiten.Image, objs *objects.Objects) {

	if l.state.HasCounted(7) {
		objs.Symbols = append(objs.Symbols, &dn{})
		l.state.Reset()
	}

	for _, sym := range objs.Symbols {
		sym.Update()
	}

	l.state.Update()

	// draw
	dg := &draw.Group{Dst: screen}

	// draw symbols
	for _, sym := range objs.Symbols {
		b := sym.Base()
		sprites.Symbol(b.IsLove).Draw(
			dg, draw.Shift(c2f(b.Pos)),
		)
	}
}

func c2f(c complex128) (float64, float64) {
	return real(c), imag(c)
}

type dn struct {
	objects.SymbolBase
	pos complex128
}

func (d *dn) Update() {
	d.pos += 1 + 2i
	d.IsLove = true
}

func (d *dn) Alpha() float64 { return 1 }
