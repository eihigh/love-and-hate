package draw

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var (
	op  = &ebiten.DrawImageOptions{}
	nop = &ebiten.DrawImageOptions{}
)

type Group struct {
	Dst *ebiten.Image
	// ColorM ebiten.ColorM
	Mode   ebiten.CompositeMode
	geoM   ebiten.GeoM
	colorM ebiten.ColorM
}

func (g *Group) Draw(src *ebiten.Image, fns ...OptionFn) {
	g.geoM.Reset()
	g.colorM.Reset()
	for _, fn := range fns {
		fn(g)
	}

	op.GeoM = g.geoM
	op.ColorM = g.colorM
	op.CompositeMode = g.Mode
	g.Dst.DrawImage(src, op)
}

func (g *Group) DrawRect(r *sio.Rect, clr color.Color) {
	if r.W < 0.5 || r.H < 0.5 {
		return
	}
	i, _ := ebiten.NewImage(int(r.W), int(r.H), ebiten.FilterDefault)
	i.Fill(clr)
	x, y := r.Pos(7)
	nop.GeoM.Translate(x, y)
	g.Dst.DrawImage(i, nop)
	nop.GeoM.Reset()
}

type OptionFn func(*Group)

func Shift(x, y float64) OptionFn {
	return func(g *Group) {
		g.geoM.Translate(x, y)
	}
}

func Paint(r, g, b, a float64) OptionFn {
	return func(gr *Group) {
		gr.colorM.Scale(r, g, b, a)
	}
}
