package draw

import "github.com/hajimehoshi/ebiten"

var op = &ebiten.DrawImageOptions{}

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
