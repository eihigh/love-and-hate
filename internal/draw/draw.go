package draw

import (
	"image/color"

	"github.com/eihigh/sio"
	"github.com/hajimehoshi/ebiten"
)

var (
	op  = &ebiten.DrawImageOptions{}
	top = &ebiten.DrawTrianglesOptions{}

	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
)

func init() {
	emptyImage.Fill(color.White)
}

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

func (gr *Group) DrawRect(re *sio.Rect, clr color.Color) {
	vs := []ebiten.Vertex{
		{
			DstX: float32(re.X),
			DstY: float32(re.Y),
		},
		{
			DstX: float32(re.X + re.W),
			DstY: float32(re.Y),
		},
		{
			DstX: float32(re.X + re.W),
			DstY: float32(re.Y + re.H),
		},
		{
			DstX: float32(re.X),
			DstY: float32(re.Y + re.H),
		},
	}

	for _, v := range vs {
		v.ColorR = 1
		v.ColorG = 1
		v.ColorB = 1
		v.ColorA = 1
	}

	indices := []uint16{
		0, 1, 3, 1, 3, 2,
	}

	gr.colorM.Reset()
	gr.colorM.Apply(clr)
	top.ColorM = gr.colorM
	gr.Dst.DrawTriangles(vs, indices, emptyImage, top)
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
